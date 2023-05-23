package middleware

import "net/http"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff
		next.ServeHTTP(w, r)
	})
}
