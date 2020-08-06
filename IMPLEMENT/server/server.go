
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
}