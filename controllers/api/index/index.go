package index

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Index(ctx *gin.Context) {
  ctx.String(http.StatusOK, "ok")
}
