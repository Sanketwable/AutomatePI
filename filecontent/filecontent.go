package filecontent

//Maincontent is a variable
var Maincontent = []byte(`package main
import ("api/server")
func main() {
	server.Run()
}`)

// GomodContent is a variable
var GomodContent = []byte(`
module api

go 1.13
`)

//ServerContent is a func
var ServerContent = []byte(
	`
	package server

import (
	"api/router"
	"fmt"
	"log"
	"net/http"
)

func Run () {
	Listen(9000)
}

func Listen(port int) {
	r := router.New()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("error is : ", err)
	}
}`)

//RouterContent is a func
var RouterContent = []byte(
	`
	package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

//New is function
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutesWithMiddlewares(r)
}`)

//RoutesContent is a func
var RoutesContent = []byte(
	`
	package routes

import (
	"api/middlewares"
	"net/http"
	"github.com/gorilla/mux"
)

// Route is a struct
type Route struct {
	Uri string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Load is  a func
func Load() []Route {
	routes := usersRoutes
	return routes
}

//SetUpRoutes is a func
func SetUpRoutes (r *mux.Router) *mux.Router {

	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
//SetUpRoutesWithMiddlewares is  a func
func SetUpRoutesWithMiddlewares (r *mux.Router) *mux.Router {

	for _, route := range Load() {
		if route.AuthRequired {
			r.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler)))).Methods(route.Method)

		} else {
			r.HandleFunc(route.Uri,middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(route.Handler))).Methods(route.Method)
		}
	}
	return r
}`)


//UserRoutesContent is a func
var UserRoutesContent = []byte(
	`
	package routes

import (
	"api/controllers"
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
}`)

//MiddlwaresContent is a func
var MiddlewaresContent = []byte(
	`package middlewares

	import (
		"api/auth"
		"api/responses"
		//"api/utils/console"
		"log"
		"net/http"
	)
	
	func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request)  {
			log.Printf("\n%s %s%s %s",r.Method, r.Host, r.RequestURI, r.Proto)
			next (w, r)
		}
	}
	
	func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request)  {
			w.Header().Set("Content-Type", "application/json")
			next (w, r)
		}
	}
	
	func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request)  {
			err := auth.TokenValid(r)
			if err != nil {
				responses.ERROR(w, http.StatusUnauthorized, err)
				return
			} 
			next (w, r)
		}
	}`)
