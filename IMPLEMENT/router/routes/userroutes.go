
	package routes

import (
	"net/http"
)

var usersRoutes = []Route {
	 {
		Uri: "/users",
		Method: http.MethodGet,
		Handler: nil,
		AuthRequired: false,
	},
	{
		Uri: "/users",
		Method: http.MethodPost ,
		Handler: nil,
		AuthRequired: false,
	},
	{
		Uri: "/users/{id}",
		Method: http.MethodGet,
		Handler: nil,
		AuthRequired: false,
	},
	{
		Uri: "/users/{id}",
		Method: http.MethodPut,
		Handler: nil,
		AuthRequired: true,
	},
	{
		Uri: "/users/{id}",
		Method: http.MethodDelete,
		Handler: nil,
		AuthRequired: true,
	},
}