package router

import "github.com/gorilla/mux"

// Generate returns a new router
func Generate() *mux.Router {
	return mux.NewRouter()
}
