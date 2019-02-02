package utils

import (
	"net/http"
	"context"
)

func WithHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(ACCESS_TOKEN)
		ctx := context.WithValue(r.Context(), ACCESS_TOKEN, accessToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}