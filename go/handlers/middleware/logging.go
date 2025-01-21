package middleware

import (
	"bytes"
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
	buffer *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	// Keep a copy of the response for logging
	rw.buffer.Write(b)
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create wrapped response writer to capture status code and response
		wrapped := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
			buffer:         &bytes.Buffer{},
		}

		// Log request
		slog.Info("Incoming request",
			"path", r.URL.Path,
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent())

		// Call the next handler
		next.ServeHTTP(wrapped, r)

		// Log response
		duration := time.Since(start)

		level := slog.LevelInfo
		attrs := []any{
			"path", r.URL.Path,
			"method", r.Method,
			"status", wrapped.status,
			"size", wrapped.size,
			"duration_ms", duration.Milliseconds(),
		}

		// Add error message for non-200 responses
		if wrapped.status >= 400 {
			level = slog.LevelWarn
			if wrapped.status >= 500 {
				level = slog.LevelError
			}

			// Add error message from response if available
			if wrapped.buffer.Len() > 0 {
				attrs = append(attrs, "error", wrapped.buffer.String())
			}
		}

		slog.Log(r.Context(), level, "Request completed", attrs...)
	}
}
