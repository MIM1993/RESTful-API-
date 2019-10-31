package router

import (
	"APISERVER/handler/sd"
	"APISERVER/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(engine *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//middlewares
	engine.Use(gin.Recovery())
	engine.Use(middleware.NoCache())
	engine.Use(middleware.Options())
	engine.Use(middleware.Secure())

	engine.Use(mw...)

	//NotFound 404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router")
	})

	//the health check handler
	svcd := engine.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DeskCheck)
		svcd.GET("/cpu", sd.CpuCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	//

	return engine
}
