package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index controller
func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
