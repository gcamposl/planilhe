package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router represents all routes from API
type Router struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Generate returns a new router
func Generate() *mux.Router {
	return mux.NewRouter()
}
