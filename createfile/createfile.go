package createfile

import (
	"AutomatePI/filecontent"
	"fmt"
	"log"
	"os"
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