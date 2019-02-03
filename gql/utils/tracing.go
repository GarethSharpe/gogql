package utils

import (
	"fmt"
	"net/http"
)

func WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Tracing request for %s", r.RequestURI)
		fmt.Println()
		next.ServeHTTP(w, r)
	}
}
