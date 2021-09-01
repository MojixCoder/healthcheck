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

// TestMySQLHealthCheck performs tests for valid and invalid data
//
// Note that being unable to connect to DB is not an exception.
// it's an ok result, but we couldn't connect to DB.
// for example trying to connect to a DB which is down.
//
// Exception raises when we have validation errors, etc.
func TestMySQLHealthCheck(t *testing.T) {
	config.Init()
	db.Init()
	router := server.NewRouter()

	values := []map[string]string{
		{"username": "root", "password": "", "host": "127.0.0.1", "port": "3306", "DBName": "YOUR_DB"},        // Valid data
		{"username": "root", "password": "", "host": "invalid hostname", "port": "3306", "DBName": "YOUR_DB"}, // Invalid data
	}

	for i, value := range values {
		jsonValue, _ := json.Marshal(value)
		res := performRequest(http.MethodPost, "/v1/mysql", router, bytes.NewBuffer(jsonValue))
		if i == 0 {
			assert.Equal(t, http.StatusOK, res.Code)
		} else if i == 1 {
			assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		}
	}
}
