package middleware

import (
	"net/http"
)

func New(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		next.ServeHTTP(w, req)
	})
}
