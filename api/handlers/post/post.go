package post

import (
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"net/http"
)

type Params struct {
	Cfg      config.Config
	Logger   logger.Logger
	Services serviceconn.ServiceController
}

type PostHandler interface {
	GetPost(rw http.ResponseWriter, r *http.Request)
	UpdatePost(rw http.ResponseWriter, r *http.Request)
	DeletePost(rw http.ResponseWriter, r *http.Request)
	GetPosts(rw http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	cfg      config.Config
	services serviceconn.ServiceController
	logger.Logger
}

func NewPostHandler(params Params) PostHandler {
	return &postHandler{
		cfg:      params.Cfg,
		Logger:   params.Logger,
		services: params.Services,
	}
}
