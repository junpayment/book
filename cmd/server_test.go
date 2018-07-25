package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetEngine ... check start server
func TestGetEngine(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ts := httptest.NewServer(GetEngine())
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/_ah/health", nil)
	cli := ts.Client()
	res, err := cli.Do(req)

	assert.Equal(t, err, nil)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	buf, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "ok", string(buf))
}
