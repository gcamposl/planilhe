package routers

import (
	"api/internal/controllers"
	"net/http"
)

var transactionRoutes = []Router{
	{
		URI:         "/transactions",
		Method:      http.MethodPost,
		Function:    controllers.CreateTransaction,
		RequireAuth: true,
	},
	{
		URI:         "/transactions/{userID}",
		Method:      http.MethodGet,
		Function:    controllers.GetTransactions,
		RequireAuth: true,
	},
}
