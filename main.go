package main

import (
	"APISERVER/router"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//create the gin engine
	engine := gin.New()

	//gin middlewares
	middlewares := []gin.HandlerFunc{}

	//router
	router.Load(engine, middlewares...)

	//ping the server to make sure the router is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up. err:", err)
		}
		log.Println("the router has been deploy successfully")
	}()

	//service
	log.Println("Start tio listening the requests on http server address %s", ":8000")
	//http.ListenAndServe err
	log.Println("http.ListenAndServe err" + http.ListenAndServe(":8000", engine).Error())

	//exit signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	//TODO: 资源回收
}

//ping function
func pingServer() error {
	for i := 0; i < 2; i++ {
		//Ping the server by Get method request `/sd/health`
		resp, err := http.Get("http://127.0.0.1:8000/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		//Sleep for a second to connect the next ping
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("cannot connect the server !")
}
