package api

import (
	"api-gateway-iman/api/handlers"
	"api-gateway-iman/api/handlers/middleware"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"

	"github.com/gorilla/mux"
)

type Params struct {
	Logger   logger.Logger
	Cfg      config.Config
	Services serviceconn.ServiceController
}

func NewRouter(params Params) *mux.Router {

	r := mux.NewRouter()
	handlers := handlers.NewHandlers(handlers.Params{Logger: params.Logger, Cfg: params.Cfg, Services: params.Services})

	r.HandleFunc("/post/get", middleware.ApplyMiddleware(handlers.PostHandler.GetPost, handlers.Middleware.LogRequest))
	r.HandleFunc("/post/update", middleware.ApplyMiddleware(handlers.PostHandler.UpdatePost, handlers.Middleware.LogRequest))
	r.HandleFunc("/post/delete/{id}", middleware.ApplyMiddleware(handlers.PostHandler.DeletePost, handlers.Middleware.LogRequest))
	r.HandleFunc("/post/list", middleware.ApplyMiddleware(handlers.PostHandler.GetPosts, handlers.Middleware.LogRequest))

	r.HandleFunc("/load/posts", middleware.ApplyMiddleware(handlers.PostLoaderHandler.LoadPosts, handlers.Middleware.LogRequest))
	r.HandleFunc("/load/status/{id}", middleware.ApplyMiddleware(handlers.PostLoaderHandler.GetLoadingStatus, handlers.Middleware.LogRequest))

	return r
}
