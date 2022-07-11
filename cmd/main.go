package main

import (
	"api-gateway-iman/api"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"fmt"
	"net/http"
)

func main() {

	cfg := config.NewConfig()
	logger := logger.New(cfg.GetString("app.log.level"), cfg.GetString("app.name"))
	fmt.Println("app port :", cfg.GetString("app.port"))
	services := serviceconn.NewServiceController(cfg)

	router := api.NewRouter(api.Params{Cfg: cfg, Logger: logger, Services: services})
	logger.Info(fmt.Sprintf("Server has been started on port: %s", cfg.GetString("app.port")))
	panic(http.ListenAndServe(cfg.GetString("app.port"), router))
}
