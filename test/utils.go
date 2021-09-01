package test

import (
	"bytes"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// performRequest performs request to a URL
func performRequest(method, target string, router *gin.Engine, body *bytes.Buffer) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, target, body)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
