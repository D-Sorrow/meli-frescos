package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/melisource/fury_go-core/pkg/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInterface interface {
	On(methodName string, arguments ...interface{}) *mock.Call
}

type TestGenericStruct[T interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}

type HandlerTestStruct struct {
	Name                    string
	ServiceMethod           string
	HttpPath                string
	HttpMethod              string
	HttpBody                []byte
	MockParams              []interface{}
	MockResponse            interface{}
	MockError               error
	ExpectedOutput          interface{}
	IsExpectedOutputMapType bool
	ExpectedStatusCode      int
	ExpectedCalls           int
}

type ServiceTestStruct struct {
	Name             string
	RepositoryMethod string
	ServiceParams    []interface{}
	MockParams       []interface{}
	MockResponse     interface{}
	MockError        error
	ExpectedOutput   interface{}
	ExpectedErr      error
}

func assertHandlerResponse[T interface{}](
	t *testing.T,
	responseDTO TestGenericStruct[T],
	responseStatusCode int,
	expectedStatusCode int,
	expectedOutput interface{},
) {
	t.Helper()

	assert.Equal(t, expectedStatusCode, responseStatusCode, "HTTP Status Code mismatch")

	if expectedOutput != nil {
		expectedDTO, ok := expectedOutput.(TestGenericStruct[T])
		if !ok {
			t.Fatal("Failed to convert expectedOutput to BuyerGenericStruct")
		}

		assert.Equal(t, expectedDTO.Code, responseDTO.Code, "ResponseDTO Code mismatch")
		assert.Equal(t, expectedDTO.Msg, responseDTO.Msg, "ResponseDTO Message mismatch")
		assert.Equal(t, expectedDTO.Data, responseDTO.Data, "ResponseDTO Data mismatch")
	}
}

func getGenericStruct[T interface{}](
	t *testing.T,
	body io.ReadCloser,
) (responseDTOGenericStruct T) {
	t.Helper()

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		t.Fatal("Body reading error:", err)
	}
	defer body.Close()

	if err := json.Unmarshal(bodyBytes, &responseDTOGenericStruct); err != nil {
		t.Fatal("JSON parsing error:", err)
	}

	return
}

func PtrStr(s string) *string {
	return &s
}

func PtrInt(i int) *int {
	return &i
}

func PtrIntNil() *int {
	return nil
}

func InitServiceMock[T MockInterface](
	t *testing.T,
	test HandlerTestStruct,
	mockService *T,
) {
	t.Helper()

	if test.ExpectedCalls > 0 {
		if test.MockResponse != nil {
			(*mockService).On(test.ServiceMethod, test.MockParams...).
				Return(test.MockResponse, test.MockError).Times(test.ExpectedCalls)
		} else {
			(*mockService).On(test.ServiceMethod, test.MockParams...).Return(test.MockError).Times(test.ExpectedCalls)
		}
	}
}

func InitRepositoryMock[T MockInterface](
	t *testing.T,
	test ServiceTestStruct,
	mockRepository *T,
) {
	t.Helper()

	if test.MockResponse != nil {
		(*mockRepository).On(test.RepositoryMethod, test.MockParams...).
			Return(test.MockResponse, test.MockError)
	} else {
		(*mockRepository).On(test.RepositoryMethod, test.MockParams...).Return(test.MockError)
	}
}

func GetResult(
	t *testing.T,
	rt *web.Router,
	path string,
	httpMethod string,
	httpBody []byte,
) (result *http.Response) {
	t.Helper()

	var req *http.Request

	if httpBody != nil {
		req = httptest.NewRequest(httpMethod, path, bytes.NewBuffer(httpBody))
	} else {
		req = httptest.NewRequest(httpMethod, path, nil)
	}

	rec := httptest.NewRecorder()
	rt.ServeHTTP(rec, req)
	result = rec.Result()
	return
}

func CheckHandlerResponse[T interface{}](
	t *testing.T,
	test HandlerTestStruct,
	result *http.Response,
) {
	if test.IsExpectedOutputMapType {
		responseDTO := getGenericStruct[TestGenericStruct[map[string]interface{}]](
			t,
			result.Body,
		)
		assertHandlerResponse(
			t,
			responseDTO,
			result.StatusCode,
			test.ExpectedStatusCode,
			test.ExpectedOutput,
		)
	} else {
		responseDTO := getGenericStruct[TestGenericStruct[T]](
			t,
			result.Body,
		)
		assertHandlerResponse(
			t,
			responseDTO,
			result.StatusCode,
			test.ExpectedStatusCode,
			test.ExpectedOutput,
		)
	}
}

func CheckServiceResponse(
	t *testing.T,
	responseErr error,
	expectedErr error,
	responseOutput interface{},
	expectedOutput interface{},
) {
	t.Helper()

	assert.True(t, errors.Is(responseErr, expectedErr), "Error mismatch")
	assert.Equal(t, responseOutput, expectedOutput, "Output mismatch")
}
