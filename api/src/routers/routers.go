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

// Configure configure all routes
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Method, route.Function).Methods(route.Method)
	}

	return r
}