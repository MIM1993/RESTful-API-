package main

import (
	"APISERVER/config"
	"APISERVER/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	for {
		fmt.Println(viper.GetString("runmode"))
		time.Sleep(4 * time.Second)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	//create the gin engine
	engine := gin.New()

	//gin middlewares
	middlewares := []gin.HandlerFunc{}

	//router
	router.Load(engine, middlewares...)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	//service
	log.Printf("Start tio listening the requests on http server address %s", viper.GetString("addr"))

	//http.ListenAndServe err
	log.Println("http.ListenAndServe err" + http.ListenAndServe(viper.GetString("addr"), engine).Error())

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
