package index

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/junpayment/book/models"
	)

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/", Index) // set handler

	// request
	req, _ := http.NewRequest("GET", "/" + "?query=東野圭吾" , nil)
	r.ServeHTTP(w, req)

	// assert
	b := &models.Book{}
	err := json.Unmarshal(w.Body.Bytes(), b)
	assert.IsType(t, nil, err)
}

func TestIndex_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/", Index) // set handler

	// request
	req, _ := http.NewRequest("GET", "/" + "?query=東野圭吾" , nil)
	r.ServeHTTP(w, req)

	// assert
	b := &models.Book{}
	err := json.Unmarshal(w.Body.Bytes(), b)
	assert.IsType(t, nil, err)
}
