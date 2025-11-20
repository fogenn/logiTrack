package httpapi

import (
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"logiTrack/internal/logger"
)

const apiKey = "4b1e1a69-3b1e-4f5b-8e1a-693b1e4f5b8e"

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logger.Log.WithFields(
				logger.Fields{
					"method": r.Method,
					"url":    r.URL.Path,
					"remote": r.RemoteAddr,
				},
			).Info("Started handling request")

			wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)

			logger.Log.WithFields(
				logger.Fields{
					"method":      r.Method,
					"url":         r.URL.Path,
					"status_code": wrapped.statusCode,
					"duration":    time.Since(start),
				},
			).Info("Finished handling request")
		},
	)
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Log.WithFields(
						logger.Fields{
							"method": r.Method,
							"url":    r.URL.Path,
							"error":  err,
							"stack":  string(debug.Stack()),
						},
					).Error("Panic occurred")

					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Internal Server Error"))
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Log.WithFields(
					logger.Fields{
						"method": r.Method,
						"url":    r.URL.Path,
						"error":  "Missing Authorization header",
					},
				).Warn("Unauthorized access attempt")

				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			parsToken := strings.TrimPrefix(authHeader, "Bearer ")
			if parsToken == authHeader {
				logger.Log.WithFields(
					logger.Fields{
						"method": r.Method,
						"url":    r.URL.Path,
						"error":  "authorization header format invalid",
					},
				).Warn("Invalid authorization format")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			if parsToken != apiKey {
				logger.Log.WithFields(
					logger.Fields{
						"method": r.Method,
						"url":    r.URL.Path,
						"error":  "Invalid API Key",
					},
				).Warn("Invalid API Key")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}

			logger.Log.WithFields(
				logger.Fields{
					"method": r.Method,
					"url":    r.URL.Path,
				},
			).Info("Authorized access")

			next.ServeHTTP(w, r)
		},
	)
}
