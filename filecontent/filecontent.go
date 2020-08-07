package filecontent

//Maincontent is a variable
var Maincontent = []byte(`package main
import (
	"api/server"
	"fmt")
func main() {
	fmt.Println("server starting")
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

	//AuthContent is a variable
	var AuthContent = []byte(
		`package auth

		import (
			"api/database"
			"api/models"
			"api/security"
			"api/utils/channels"
			"github.com/jinzhu/gorm"
		)
		
		//SignIn is func
		func SignIn(email, password string) (string, error) {
			user := models.User{}
			var err error
			var db *gorm.DB
			done := make(chan bool)
		
			go func(ch chan<- bool) {
				defer close(ch)
				db, err = database.Connect()
				if err != nil {
					ch <- false
					return
				}
				defer db.Close()
		
		
		
				err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
				if err != nil {
					ch <- false
					return
				}
				err = security.VerifyPassword(user.Password, password)
				if err != nil {
					ch <- false
					return
				}
				ch <- true
			}(done)
		
			if channels.OK(done) {
				return CreateToken(user.ID)
			}
			return "", err
		
		}`)
	//TokenContent is a variable
	var TokenContent = []byte(
		`package auth

		import (
			"api/config"
			"api/utils/console"
			"fmt"
			"net/http"
			//"reflect"
			"strconv"
		
			//"strconv"
			"strings"
			"time"
		
			"github.com/dgrijalva/jwt-go"
		)
		
		
		func CreateToken (user_id uint32) (string, error) {
			claims := jwt.MapClaims{}
			claims["authorized"] = true
			claims["user_id"] = user_id
			claims["exp"] = time.Now().Add(time.Hour * 100).Unix()
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			return token.SignedString(config.SECRETKEY)
		}
		
		func TokenValid(r *http.Request ) error {
			tokenString := ExtractToken(r)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return config.SECRETKEY, nil
			})
		
			if err != nil {
				return err
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				console.Pretty(claims)
			}
			return nil
		}
		
		func ExtractToken(r *http.Request) string {
			keys := r.URL.Query()
			token := keys.Get("token")
			if token != "" {
				return token
			}
			bearerToken := r.Header.Get("Authorization")
		
			if len(strings.Split(bearerToken, " ")) == 2 {
				return strings.Split(bearerToken, " ")[1]
			}
			return ""
		
		}
		
		func ExtractTokenID(r *http.Request) (uint32, error) {
			tokenString := ExtractToken(r)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return config.SECRETKEY, nil
			})
		
			if err != nil {
				return 0, err
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
				if err != nil {
					return 0, err
				}
					
					return uint32(uid), nil
			}
			return 0, nil
		}`)

	//DataContent is a var
	var DataContent = []byte(
		`package auto

		import "api/models"
		
		var users = []models.User{
			{
				UserName: "sanket",
				Email: "sanketwable312@gmail.com",
				Password: "SanketWable@123",
			},
		}
		
		var posts = []models.Post {
			{
				Name: "Sanket",
				AboutMe: "I m pandit",
				Age: 11,
				YearsOfExperience: 3,
				Education: "NITh cse",
				DetailsOnVidya: "No detail",
			},
		}`)

	//LoadContent is a var
	var LoadContent = []byte(
		`package auto

		import (
			"api/database"
			"api/models"
			//"api/utils/console"
			"log"
		)
		//Load is
		func Load()  {
			db, err := database.Connect()
			if err != nil {
				log.Fatal("this is an error :", err)
			}
			defer db.Close()
		
			// err = db.Debug().DropTableIfExists(&models.User{}, &models.Post{}, &models.ExpertiseService{}, &models.PujaService{}, &models.OtherserviceService{}, models.PujaServiceDuration{}, models.PujaServicePrice{}, models.Availability{}, models.DateTime{}, models.Verification{}).Error
			// if err != nil {
			// 	log.Fatal("this is an error :", err)
			// }
		
			err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.ExpertiseService{}, &models.PujaService{}, &models.OtherserviceService{}, models.PujaServiceDuration{}, models.PujaServicePrice{}, models.Availability{}, models.DateTime{}, models.Verification{}, models.BookPuja{}, models.FrontPageLoader{}, models.PujaServiceVideo{}).Error
			if err != nil {
				 log.Fatal("error occured : ", err )
			 }
		
			err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id","users(id)","cascade","cascade").Error
			if err != nil {
				log.Fatal("error occured : ", err )
			}
		
			/*
			 for i := range users {
				err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
				if err != nil {
					log.Fatal("error occured : ", err )
				}
				posts[i].AuthorID = users[i].ID
				err = db.Debug().Model(&models.User{}).Create(&posts[i]).Error
				if err != nil {
					log.Fatal("error occured : ", err )
				}
				
			 }*/
		
		}`)
var DBContent = []byte(
	`package database

	import (
		"api/config"
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
	)
	
	func Connect() (*gorm.DB, error) {
		db, err := gorm.Open(config.DBDRIVER, config.DBURL)
		if err != nil {
			return nil, err
		}
		return db, nil
	
	}`)

var ConfigContent = []byte(
`
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//PORT is port to which it has to be connected
	PORT = 0
	SECRETKEY []byte
	DBURL = ""
	DBDRIVER =""
	STOREURL []byte
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error is : ", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT")) 

	if err != nil {
		PORT = 8080
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),os.Getenv("DB_NAME"))

	SECRETKEY = []byte(os.Getenv("API_SECRET"))

	STOREURL = []byte(os.Getenv("STORE_URL"))

	
}`)

var UserControllerContent = []byte(
	`
	package controllers

import (
	"api/auth"
	"api/database"
	"api/models"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetUsers is a func
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		users, err := usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

//CreateUser is a func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	

	body, err :=ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)

}

//GetUser is a func
func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		user, err := usersRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

//UpdateUser is a func
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err :=ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if tokenID != uint32(uid){
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

 
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		rows, err := usersRepository.Update(uint32(uid), user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}

//DeleteUser is a func
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryUsersCRUD(db)

	func(usersRepository repository.UserRepository) {
		_, err := usersRepository.Delete(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON(w, http.StatusNoContent , "")
	}(repo)
}`)
var UserModelContent = []byte(
	`package models

	import (
		"api/security"
		"errors"
		"html"
		"log"
		"strings"
		"time"
	
		"github.com/badoux/checkmail"
		//"github.com/sqs/goreturns/returns"
	)
	
	//User is a struct
	type User struct {
		ID        uint32    replace gorm:"primary_key;auto_increment" json:"id" replace
		UserName  string    replace gorm:"size:20;not null;unique" json:"username" replace
		Email     string    replace gorm:"size:50;not null;unique" json:"email" replace
		Password  string    replace gorm:"size:60;not null" json:"password" replace
		CreatedAt time.Time replace gorm:"" json:"created_at" replace
		UpdatedAt time.Time replace gorm:"" json:"updated_at" replace
		
	}
	
	//BeforeSave is a func
	func (u *User) BeforeSave() error {
		hashedPassword, err  := security.Hash(u.Password)
		if err != nil {
			log.Fatal("error comes : ", err)
		}
	
		u.Password = string(hashedPassword)
		return nil
	}
	
	//Prepare is a struct
	func (u *User) Prepare() {
		u.ID = 0
		u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
		u.Email = html.EscapeString(strings.TrimSpace(u.Email))
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	
	}
	
	//Validate is a func
	func (u *User) Validate(action string) error {
		switch action {
		case "update":
			if u.UserName == "" {
				return errors.New("Required UserName")
			}
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			return nil
		case "login":
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			return nil
		case "signup" :
			if u.Email == "" {
				return errors.New("Required email")
			}
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			return nil
		default:
			if u.UserName == "" {
				return errors.New("Required username")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			return nil
		}
	}`)
var RepositoryUserContent = []byte(
	`package repository

	import "api/models"
	
	type UserRepository interface {
		Save(models.User) (models.User, error)
		FindAll() ([]models.User, error)
		FindById(uint32) (models.User, error)
		Update(uint32, models.User) (int64, error)
		Delete(uint32) (int64, error)
	}`)

var RepositoryUserCRUDContent = []byte(
	`package crud

	import (
		"api/models"
		"api/utils/channels"
		"errors"
	
		"github.com/jinzhu/gorm"
	)
	
	type repositoryUsersCRUD struct {
		db *gorm.DB
	}
	//NewRepositoryUsersCRUD is func
	func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD{
		return &repositoryUsersCRUD{db}
	}
	
	func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
		var err error
		done := make(chan bool) 
		go func(ch chan<- bool) {
			err = r.db.Debug().Model(models.User{}).Create(&user).Error
			if err != nil {
				ch <- false
				return
			}
			ch <- true
		}(done)
	
		if channels.OK(done) {
			return user, nil
		}
		return models.User{}, err
	}
	
	func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
		var err error
	
		users := []models.User{}
	
		done := make(chan bool) 
		go func(ch chan<- bool) {
			defer close(ch)
			err = r.db.Debug().Model(models.User{}).Limit(100).Find(&users).Error
			if err != nil {
				ch <- false
				return
			}
			ch <- true
		}(done)
	
		if channels.OK(done) {
			return users, nil
		}
		return nil, err
	}
	
	func (r *repositoryUsersCRUD) FindById(uid uint32) (models.User, error) {
		var err error
	
		user := models.User{}
	
		done := make(chan bool) 
		go func(ch chan<- bool) {
			defer close(ch)
	
			err = r.db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
			if err != nil {
				ch <- false
				return
			}
			ch <- true
		}(done)
	
		if channels.OK(done) {
			return user, nil
		}
		if gorm.IsRecordNotFoundError(err) {
			return models.User{ }, errors.New("user not found")
		}
		return models.User{ }, err
	}
	
	func (r *repositoryUsersCRUD) Update(uid uint32, user models.User) (int64, error) {
		var rs *gorm.DB
	 
	
		done := make(chan bool) 
		go func(ch chan<- bool) {
			defer close(ch)
			rs = r.db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(user)
			ch <- true
		}(done)
	
		if channels.OK(done) {
			if rs.Error != nil {
				return 0, rs.Error
			}
			return rs.RowsAffected, nil
		}
		return 0, rs.Error
	}
	
	func (r *repositoryUsersCRUD) Delete(uid uint32) (int64, error) {
		var rs *gorm.DB 
		done := make(chan bool) 
		go func(ch chan<- bool) {
			defer close(ch)
			rs = r.db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
			ch <- true
		}(done)
	
		if channels.OK(done) {
			if rs.Error != nil {
				return 0, rs.Error
			}
			return rs.RowsAffected, nil
		}
		return 0, rs.Error
	}`)
var JSONContent = []byte(
	`package responses

	import (
		"encoding/json"
		"fmt"
		"net/http"
	)
	
	func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
		w.WriteHeader(statusCode)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	
	}
	
	func ERROR(w http.ResponseWriter, statusCode int, err error) {
		if err != nil {
			JSON(w, statusCode, struct {
				Error string replace json:"error" replace
			}{
				Error: err.Error(),
			})
		}
		JSON(w, http.StatusBadRequest, nil)
	}`)

var SecurityContent = []byte(
	`package security

	import "golang.org/x/crypto/bcrypt"
	//Hash is used to produce hash
	func Hash(password string) ([]byte, error) {
		return bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	
	}
	//VerifyPassword is used to verify the pasoword hash
	func VerifyPassword(hashedPassword, password string) error {
		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}`)

var ConsoleContent = []byte(
	`
	package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal("error  is : ", err)
		return
	}
	fmt.Println(string(b))

}`)

var ChannelsContent = []byte(
	`package channels

	func OK (done chan bool) bool {
		select {
		case ok := <-done:
			if ok {
				return true
			}
		}
		return false
	}`)

var DockerContent = []byte(
	`FROM golang:latest

	RUN mkdir /appserver/
	ADD . /appserver/
	WORKDIR /appserver/
	RUN go build -o main .
	CMD ["/appserver/main"]`)

var ENVContent = []byte(
	`API_PORT=9000
	API_SECRET=snf78y34jnh9734jhgf894hf
	
	#DATABASE CONFIG
	DB_DRIVER=mysql
	DB_USER=sanket
	DB_PASS=Sanket@123
	DB_NAME=panditji
	DB_HOST=db.pdjt.prod
	
	#Store files folder
	STORE_URL=/Users/sanketwable/desktop`)
