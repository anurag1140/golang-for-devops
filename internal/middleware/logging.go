package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (l *loggingResponseWriter) WriteHeader(status int) {

	l.status = status

	l.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		start := time.Now()

		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)

		slog.Info(
			"HTTP Request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", lrw.status,
			"duration", duration,
		)
	})
}
