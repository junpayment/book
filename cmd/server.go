package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// health check
	router.GET("/_ah/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// web
	web := router.Group("/web")
	{
		web.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	// api
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	router.Run(":8080")
}
