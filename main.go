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
	os.MkdirAll("IMPLEMENT/auth", 0700)
	os.MkdirAll("IMPLEMENT/auto", 0700)
	os.MkdirAll("IMPLEMENT/database", 0700)
	os.MkdirAll("IMPLEMENT/config", 0700)
	os.MkdirAll("IMPLEMENT/controllers", 0700)
	os.MkdirAll("IMPLEMENT/models", 0700)
	os.MkdirAll("IMPLEMENT/repository/CRUD", 0700)
	os.MkdirAll("IMPLEMENT/responses", 0700)
	os.MkdirAll("IMPLEMENT/security", 0700)
	os.MkdirAll("IMPLEMENT/utils/channels", 0700)
	os.MkdirAll("IMPLEMENT/utils/console", 0700)

	getauto.GetModels()
	createfile.CreateMain()
	createfile.Creategomod()
	createfile.CreateServer()
	createfile.CreateRouter()
	createfile.CreateRoutes()
	createfile.CreateUserRoutes()
	createfile.CreateMiddlewares()
	createfile.CreateAuth()
	createfile.CreateToken()
	createfile.CreateData()
	createfile.CreateLoad()
	createfile.CreateDB()
	createfile.CreateConfig()
	createfile.CreateUsersController()
	createfile.CreateUserModel()
	createfile.CreateRepositoryUsers()
	createfile.CreateRepositoryUsersCrud()
	createfile.CreateJson()
	createfile.CreateSecurity()
	createfile.CreateConsole()
	createfile.CreateChannels()
	createfile.CreateDockerfile()
	createfile.CreateENV()



	

}
