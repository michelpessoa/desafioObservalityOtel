package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const (
	RequestIDKey contextKey = "X-Request-ID"
)

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
		w.Header().Set(string(RequestIDKey), requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(RequestIDKey).(string)
}
