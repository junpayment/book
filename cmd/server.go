package main

import (
  "github.com/gin-gonic/gin"
  "github.com/junpayment/book/controllers/api/index"
  "net/http"
)

func main() {
  GetEngine().Run(":8080")
}

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
