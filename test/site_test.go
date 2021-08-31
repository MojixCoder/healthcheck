package test

import (
	"bytes"
	"encoding/json"
	"github.com/MojixCoder/healthcheck/config"
	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestSiteHealthCheck perform tests for valid and invalid URL
func TestSiteHealthCheck(t *testing.T) {
	config.Init()
	db.Init()
	router := server.NewRouter()

	values := []map[string]string{
		{"url": "https://stackoverflow.com/questions/35362459/golang-create-a-slice-of-maps"}, // This is a valid URL
		{"url": "this is an invalid url"}, // This an invalid URL
	}

	for i, value := range values {
		jsonValue, _ := json.Marshal(value)
		res := performRequest("POST", "/v1/website", router, bytes.NewBuffer(jsonValue))
		if i == 0 {
			assert.Equal(t, http.StatusOK, res.Code)
		} else if i == 1 {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	}
}

// performRequest performs request to a URL
func performRequest(method, target string, router *gin.Engine, body *bytes.Buffer) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, body)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
