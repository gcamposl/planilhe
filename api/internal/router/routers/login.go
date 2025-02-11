package routers

import "net/http"

var loginRoute = Router{
	URI:    "/login",
	Method: http.MethodPost,
	Function: func(w http.ResponseWriter, r *http.Request) {

	},
	RequireAuth: false,
}
