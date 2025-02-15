package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ValidatePatchRequestBody(r *http.Request) (data map[string]interface{}, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &data); err != nil {
		return
	}

	if err = validateJsonKeys(data); err != nil {
		return nil, err
	}

	if err = validateValuesJsonType(data); err != nil {
		return nil, err
	}

	return
}

func validateJsonKeys(data map[string]interface{}) (err error) {
	keys := map[string]struct{}{
		"address":             {},
		"telephone":           {},
		"warehouse_code":      {},
		"minimun_capacity":    {},
		"minimun_temperature": {},
		"locality_id":         {},
	}

	for key := range data {
		if _, find := keys[key]; !find {
			err = fmt.Errorf("wrong key %s", key)
			return
		}
	}

	return
}

func validateValuesJsonType(data map[string]interface{}) (err error) {

	for key, value := range data {
		switch key {
		case "warehouse_code":
			if _, ok := value.(string); !ok {
				err = errors.New("warehouse code must be a string")
			}
		case "address":
			if _, ok := value.(string); !ok {
				err = errors.New("address must be a string")
			}
		case "telephone":
			if _, ok := value.(string); !ok {
				err = errors.New("telophone must be a string")
			}
		case "minimun_capacity":
			if _, ok := value.(float64); !ok {
				err = errors.New("minimun capacity must be a int")
			}
		case "minimun_temperature":
			if _, ok := value.(float64); !ok {
				err = errors.New("minimun temperature must be a int")
			}
		case "locality_id":
			if _, ok := value.(float64); !ok {
				err = errors.New("locality id must be a int")
			}
		}
	}
	return
}

func ValidatePatchValues(data map[string]interface{}) error {
	for key, value := range data {
		switch key {
		case "warehouse_code":
			if value == "" {
				return errors.New("warehouse code can not be empty")
			}
		case "address":
			if value == "" {
				return errors.New("address can not be empty")
			}
		case "telephone":
			if value == "" {
				return errors.New("telephone can not be empty")
			}
		case "minimun_capacity":
			if val, ok := value.(float64); ok {
				if val < 0 {
					return errors.New("capacity can not be less than cero")
				}
			}
		case "minimun_temperature":
			if val, ok := value.(float64); ok {
				if val < -18 || val > 15 {
					return errors.New("temperatura must be between -18 and 15 degrees")
				}
			}
		case "locality_id":
			if val, ok := value.(float64); ok {
				if val < 1 {
					return errors.New("locality id can not be less than 1")
				}
			}
		}
	}
	return nil
}
