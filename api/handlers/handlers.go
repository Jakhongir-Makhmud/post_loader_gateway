package handlers

import (
	"api-gateway-iman/api/handlers/middleware"
	"api-gateway-iman/api/handlers/post"
	postloader "api-gateway-iman/api/handlers/postLoader"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
)

type Params struct {
	Logger   logger.Logger
	Cfg      config.Config
	Services serviceconn.ServiceController
}

type Handlers struct {
	PostHandler       post.PostHandler
	PostLoaderHandler postloader.PostLoaderHandler
	Middleware middleware.Middleware
}

func NewHandlers(params Params) Handlers {

	postHandler := post.NewPostHandler(post.Params{
		Cfg: params.Cfg, 
		Logger: params.Logger, 
		Services: params.Services})

	postLoaderHandler := postloader.NewPostLoader(postloader.Params{
		Cfg: params.Cfg, 
		Logger: params.Logger, 
		Services: params.Services})

	middleware := middleware.NewMiddleware(params.Logger)
	
	return Handlers{postHandler, postLoaderHandler, middleware}
}
