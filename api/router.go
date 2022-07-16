package api

import (
	"api-gateway-iman/api/handlers"
	"api-gateway-iman/api/handlers/middleware"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"fmt"
	"net/http"

	_ "api-gateway-iman/api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Params struct {
	Logger   logger.Logger
	Cfg      config.Config
	Services serviceconn.ServiceController
}

// @title           Post api
// @version         1.0
// @description     This is a post server

// @contact.name   Jakhongir
// @contact.email  anorboev.jahongir8007@gmail.com
// @host      localhost:9003
// @BasePath /

func NewRouter(params Params) *mux.Router {

	r := mux.NewRouter()
	handlers := handlers.NewHandlers(handlers.Params{Logger: params.Logger, Cfg: params.Cfg, Services: params.Services})

	r.HandleFunc("/post/get/{id}", middleware.ApplyMiddleware(handlers.PostHandler.GetPost, handlers.Middleware.LogRequest)).Methods("GET")
	r.HandleFunc("/post/update", middleware.ApplyMiddleware(handlers.PostHandler.UpdatePost, handlers.Middleware.LogRequest)).Methods("PUT")
	r.HandleFunc("/post/delete/{id}", middleware.ApplyMiddleware(handlers.PostHandler.DeletePost, handlers.Middleware.LogRequest)).Methods("DELETE")
	r.HandleFunc("/post/list", middleware.ApplyMiddleware(handlers.PostHandler.GetPosts, handlers.Middleware.LogRequest)).Methods("POST")

	r.HandleFunc("/load/posts", middleware.ApplyMiddleware(handlers.PostLoaderHandler.LoadPosts, handlers.Middleware.LogRequest)).Methods("POST")
	r.HandleFunc("/load/status/{id}", middleware.ApplyMiddleware(handlers.PostLoaderHandler.GetLoadingStatus, handlers.Middleware.LogRequest)).Methods("GET")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", params.Cfg.GetString("app.port"))), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	return r
}
