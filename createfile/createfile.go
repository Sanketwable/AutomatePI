package createfile

import (
	"AutomatePI/filecontent"
	"fmt"
	"log"
	"os"
	"bytes"
	"io/ioutil"
)

//CreateModels is a func
func CreateMain() {
	mainfile, err := os.Create("IMPLEMENT/main.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	main, err := mainfile.Write(filecontent.Maincontent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", main)
	mainfile.Close()

}

//Creategomod is a func
func Creategomod() {
	gomodfile, err := os.Create("IMPLEMENT/go.mod")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	gomod, err := gomodfile.Write(filecontent.GomodContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", gomod)
	gomodfile.Close()
}
func CreateServer() {
	serverfile, err := os.Create("IMPLEMENT/server/server.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	server, err := serverfile.Write(filecontent.ServerContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", server)
	serverfile.Close()
}
func CreateRouter() {
	routerfile, err := os.Create("IMPLEMENT/router/router.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	router, err := routerfile.Write(filecontent.RouterContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", router)
	routerfile.Close()
}
func CreateRoutes() {
	routesfile, err := os.Create("IMPLEMENT/router/routes/routes.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	routes, err := routesfile.Write(filecontent.RoutesContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", routes)
	routesfile.Close()
}
func CreateUserRoutes() {
	userroutesfile, err := os.Create("IMPLEMENT/router/routes/userroutes.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	userroutes, err := userroutesfile.Write(filecontent.UserRoutesContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", userroutes)
	userroutesfile.Close()
}
func CreateMiddlewares() {
	middlewaresfile, err := os.Create("IMPLEMENT/middlewares/middlewares.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	middlewares, err := middlewaresfile.Write(filecontent.MiddlewaresContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", middlewares)
	middlewaresfile.Close()
}
func CreateAuth() {
	authfile, err := os.Create("IMPLEMENT/auth/auth.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	auth, err := authfile.Write(filecontent.AuthContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", auth)
	authfile.Close()
}
func CreateToken() {
	tokenfile, err := os.Create("IMPLEMENT/auth/token.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	token, err := tokenfile.Write(filecontent.TokenContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", token)
	tokenfile.Close()
}
func CreateData() {
	datafile, err := os.Create("IMPLEMENT/auto/data.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	data, err := datafile.Write(filecontent.DataContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", data)
	datafile.Close()
}
func CreateLoad() {
	loadfile, err := os.Create("IMPLEMENT/auto/load.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	load, err := loadfile.Write(filecontent.LoadContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", load)
	loadfile.Close()
}

func CreateDB() {
	dbfile, err := os.Create("IMPLEMENT/database/db.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	db, err := dbfile.Write(filecontent.DBContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", db)
	dbfile.Close()
}
func CreateConfig() {
	configfile, err := os.Create("IMPLEMENT/config/config.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	config, err := configfile.Write(filecontent.ConfigContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", config)
	configfile.Close()
}
func CreateUsersController() {
	usercontrollerfile, err := os.Create("IMPLEMENT/controllers/users_controller.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	usercontroller, err := usercontrollerfile.Write(filecontent.UserControllerContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", usercontroller)
	usercontrollerfile.Close()
}
func CreateUserModel() {
	usermodelfile, err := os.Create("IMPLEMENT/models/user.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	usermodel, err := usermodelfile.Write(filecontent.UserModelContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", usermodel)
	usermodelfile.Close()
	Replace("IMPLEMENT/models/user.go")
}
func CreateRepositoryUsers() {
	repositoryuserfile, err := os.Create("IMPLEMENT/repository/repository_users.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	repositoryuser, err := repositoryuserfile.Write(filecontent.RepositoryUserContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", repositoryuser)
	repositoryuserfile.Close()
}
func CreateRepositoryUsersCrud() {
	repositoryusercrudfile, err := os.Create("IMPLEMENT/repository/CRUD/repository_users_crud.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	repositoryusercrud, err := repositoryusercrudfile.Write(filecontent.RepositoryUserCRUDContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", repositoryusercrud)
	repositoryusercrudfile.Close()
}
func CreateJson() {
	jsonfile, err := os.Create("IMPLEMENT/responses/json.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	json, err := jsonfile.Write(filecontent.JSONContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", json)
	jsonfile.Close()

	Replace("IMPLEMENT/responses/json.go")
}
func CreateSecurity() {
	securityfile, err := os.Create("IMPLEMENT/security/security.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	security, err := securityfile.Write(filecontent.SecurityContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", security)
	securityfile.Close()
}
func CreateConsole() {
	consolefile, err := os.Create("IMPLEMENT/utils/console/console.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	console, err := consolefile.Write(filecontent.ConsoleContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", console)
	consolefile.Close()
}
func CreateChannels() {
	channelsfile, err := os.Create("IMPLEMENT/utils/channels/channels.go")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	channels, err := channelsfile.Write(filecontent.ChannelsContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", channels)
	channelsfile.Close()
}
func CreateDockerfile() {
	dockerfile, err := os.Create("IMPLEMENT/Dockerfile")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	docker, err := dockerfile.Write(filecontent.DockerContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", docker)
	dockerfile.Close()
}
func CreateENV() {
	envfile, err := os.Create("IMPLEMENT/.env")
	if err != nil {
		log.Fatal("error occured while making file: ", err)
	}

	env, err := envfile.Write(filecontent.ENVContent)
	if err != nil {
		log.Fatal("error occured while writing data to file: ", err)
	}
	fmt.Println("length = ", env)
	envfile.Close()
}

func Replace(fileName string) {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
			log.Fatal(err)
	}

	output := bytes.Replace(input, []byte("*replace*"), []byte("`"), -1)

	if err = ioutil.WriteFile(fileName, output, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
}