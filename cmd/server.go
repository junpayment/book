package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junpayment/book/controllers/api/index"
	"net/http"
	"log"
)

func main() {
	err := GetEngine().Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// GetEngine ...
func GetEngine() *gin.Engine {
	e := gin.Default()

	// health check
	e.GET("/_ah/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// web
	web := e.Group("/web")
	{
		web.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	// api
	api := e.Group("/api")
	{
		api.GET("/", index.Index)
	}

	return e
}
