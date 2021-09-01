package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MojixCoder/healthcheck/config"
	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/server"
	"github.com/stretchr/testify/assert"
)

// TestSiteHealthCheck performs tests for valid and invalid URL
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
		res := performRequest(http.MethodPost, "/v1/website", router, bytes.NewBuffer(jsonValue))
		if i == 0 {
			assert.Equal(t, http.StatusOK, res.Code)
		} else if i == 1 {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	}
}
