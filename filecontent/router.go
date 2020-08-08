package filecontent

//RouterContent is a func
var RouterContent = []byte(
	`
	package router
	
	import (
		"automatepi/router/routes"
	
		"github.com/gorilla/mux"
	)
	
	//New is function
	func New() *mux.Router {
		r := mux.NewRouter().StrictSlash(true)
		return routes.SetUpRoutesWithMiddlewares(r)
	}`)
	