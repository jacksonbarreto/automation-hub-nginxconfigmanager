package router

import (
	"automation-hub-nginxconfigmanager/internal/app/autoconfig"
	"automation-hub-nginxconfigmanager/internal/app/config"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/add-config", autoconfig.AddConfigHandler)
	mux.HandleFunc("/remove-config", autoconfig.RemoveConfigHandler)
	mux.HandleFunc("/update-config", autoconfig.UpdateConfigHandler)

	return http.StripPrefix(config.Config.BaseUrl, mux)
}
