package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	core_logger "gitlab.com/voykinEgor/gorestapi/internal/core/logger"
	core_response "gitlab.com/voykinEgor/gorestapi/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const (
	header_request_id = "X-Request-ID"
)

func RequestID() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(header_request_id)
			if len(requestId) == 0 {
				requestId = uuid.NewString()
			}

			r.Header.Set(header_request_id, requestId)
			w.Header().Set(header_request_id, requestId)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(header_request_id)
			l := log.With(
				zap.String("requestId", requestId),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := r.Context().Value("log").(*core_logger.Logger)
			responseHandler := core_response.NewHTTPResponseHandler(log, w)
			defer func() {
				if rec := recover(); rec != nil {
					responseHandler.PanicResponse(rec, "unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := r.Context().Value("log").(*core_logger.Logger)
			timeRequest := time.Now().UTC()

			responseWriter := core_response.NewResponseWriter(w)

			log.Debug(
				">>> incoming http request",
				zap.Time("time", timeRequest),
			)

			next.ServeHTTP(responseWriter, r)
			log.Debug(
				"<<< outcoming http response",
				zap.Int("responseCode", responseWriter.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(timeRequest)),
			)
		})
	}
}
