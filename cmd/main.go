package main

import (
	"api-gateway-iman/api"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.NewConfig()
	logger := logger.New(cfg.GetString("app.log.level"), cfg.GetString("app.name"))
	fmt.Println("app port :", cfg.GetString("app.port"))
	services := serviceconn.NewServiceController(cfg)

	router := api.NewRouter(api.Params{Cfg: cfg, Logger: logger, Services: services})
	logger.Info(fmt.Sprintf("Server has been started on port: %s", cfg.GetString("app.port")))

	server := http.Server{Addr: cfg.GetString("app.port"), Handler: router}

	shutDown := make(chan os.Signal, 1)

	signal.Notify(shutDown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutDown
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Fatal("gracefull shutdown failed")
		}
		logger.Info("server gracefully shutted down")
	}()

	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		logger.Info("server stoped")
	} else {
		panic(err)
	}
}
