package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/melisource/fury_go-core/pkg/web"

	"github.com/D-Sorrow/meli-frescos/internal/domain/models"
	"github.com/D-Sorrow/meli-frescos/internal/domain/service"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/db"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/repository"
)

type contextKey string

const AppErrorKey contextKey = "APP_ERR"

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func (rr *ResponseRecorder) WriteHeader(statusCode int) {
	rr.StatusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}

func (rr *ResponseRecorder) Write(b []byte) (int, error) {
	return rr.ResponseWriter.Write(b)
}

func LogErrorMiddleware(
	db *db.DataBase,
	ctx *context.Context,
) web.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			rr := NewResponseRecorder(w)
			next.ServeHTTP(rr, r)
			if err := (*ctx).Value(AppErrorKey); err != nil {
				logRepo := repository.NewLogRepository(db.Db)
				logService := service.NewLogService(logRepo)

				logAttributes := models.LogAttributes{
					Level:  models.ERROR,
					Source: r.URL.Path,
					Detail: fmt.Sprintf(
						"Error occurred with status code [%d], %v",
						rr.StatusCode,
						err,
					),
				}

				if _, err := logService.Register(logAttributes); err != nil {
					log.Println("Error al registrar el log:", err)
				}
			}
		}
	}
}
