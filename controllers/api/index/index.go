package index

import (
	"github.com/gin-gonic/gin"
	"github.com/junpayment/book/models"
	"net/http"
)

// Index controller
func Index(ctx *gin.Context) {
	s := models.NewBookSource()
	s.Query = ctx.Request.URL.Query().Get("query")
	r := models.NewBookRepository(s)
	b, err := r.GetBookItems()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, b)
}
