package web

import (
	"net/http"
)

// CORSHandler ...
func CORSHandler() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if o := r.Header.Get("Origin"); o != "" {
				w.Header().Set("Access-Control-Allow-Origin", o)
			}
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Method", r.Header.Get("Access-Control-Request-Method"))
				w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
				w.Header().Set("Access-Control-Max-Age", "86400")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
