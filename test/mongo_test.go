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

// TestMongoHealthCheck performs tests for valid and invalid MongoDB URI
func TestMongoHealthCheck(t *testing.T) {
	config.Init()
	db.Init()
	router := server.NewRouter()

	values := []map[string]string{
		{"mongoURI": "mongodb://127.0.0.1:27017"}, // Valid data
		{"mongoURI": "Invalid MongoDB URI"},       // Invalid data
	}

	for i, value := range values {
		jsonValue, _ := json.Marshal(value)
		res := performRequest(http.MethodPost, "/v1/mongo", router, bytes.NewBuffer(jsonValue))
		if i == 0 {
			assert.Equal(t, http.StatusOK, res.Code)
		} else if i == 1 {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	}
}
