package routers

import (
	"api/internal/controllers"
	"net/http"
)

var loginRoute = Router{
	URI:         "/login",
	Method:      http.MethodPost,
	Function:    controllers.Login,
	RequireAuth: false,
}
