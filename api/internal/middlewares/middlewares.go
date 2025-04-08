package middlewares

import (
	"fmt"
	"net/http"
)

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validation...")
		next.ServeHTTP(w, r)
	}
}
