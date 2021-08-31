package server

import (
	"github.com/MojixCoder/healthcheck/config"
)

// Init runs server
func Init() {
	appConfig := config.GetAppConfig()
	r := NewRouter()
	r.Run(appConfig.PORT)
}
