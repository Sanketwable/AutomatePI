package main

import (
	"AutomatePI/createfile"
	"AutomatePI/getauto"
	"os"
)

func main() {
	os.MkdirAll("IMPLEMENT/server", 0700)
	os.MkdirAll("IMPLEMENT/router/routes", 0700)
	os.MkdirAll("IMPLEMENT/middlewares", 0700)
	getauto.GetModels()
	createfile.CreateMain()
	createfile.Creategomod()
	createfile.CreateServer()
	createfile.CreateRouter()
	createfile.CreateRoutes()
	createfile.CreateUserRoutes()
	createfile.CreateMiddlewares()
	

}
