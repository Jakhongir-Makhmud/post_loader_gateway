package postloader

import (
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"net/http"
)


type Params struct {
	Cfg config.Config
	Logger logger.Logger
	Services serviceconn.ServiceController
}


type PostLoaderHandler interface {
	LoadPosts(rw http.ResponseWriter, r *http.Request)
	GetLoadingStatus(rw http.ResponseWriter, r *http.Request)
}

type postLoaderHandler struct {
	cfg config.Config
	logger logger.Logger
	services serviceconn.ServiceController
}


func NewPostLoader(params Params) PostLoaderHandler {

	return &postLoaderHandler{
		cfg: params.Cfg,
		logger: params.Logger,
		services: params.Services,
	}
}


func (h *postLoaderHandler) LoadPosts(rw http.ResponseWriter, r *http.Request) {

}

func (h *postLoaderHandler) GetLoadingStatus(rw http.ResponseWriter, r *http.Request) {

}