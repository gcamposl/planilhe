package middlewares

import (
	"fmt"
	"net/http"
)

// Logger writes request informations on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
	}
}

// Authenticate if user is logged in
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validation...")
		next(w, r)
	}
}
