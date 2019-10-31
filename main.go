package main

import (
	"APISERVER/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//create the gin engine
	engine := gin.New()

	//gin middlewares
	middlewares := []gin.HandlerFunc{}

	//router
	router.Load(engine, middlewares...)

	//service
	log.Println("Start tio listening the requests on http server address %s", ":8000")

	//http.ListenAndServe err
	log.Println("http.ListenAndServe err" + http.ListenAndServe(":8000", engine).Error())

}
