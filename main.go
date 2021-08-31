package main

import (
	"github.com/MojixCoder/healthcheck/config"
	"github.com/MojixCoder/healthcheck/db"
	"github.com/MojixCoder/healthcheck/server"
)

func main() {
	// Set up application config
	config.Init()

	// Connect to database
	db.Init()

	// Run application server
	server.Init()
}
