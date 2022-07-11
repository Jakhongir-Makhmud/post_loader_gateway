package serviceconn

import (
	pbl "api-gateway-iman/genproto/post_loader_service"
	pbp "api-gateway-iman/genproto/post_service"
	"api-gateway-iman/pkg/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceController interface {
	PostService() pbp.PostServiceClient
	PostLoaderService() pbl.PostLoaderServiceClient
}

type services struct {
	conns map[string]interface{}
}

func NewServiceController(cfg config.Config) ServiceController {

	postConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d",
			cfg.GetString("app.services.post_service.host"),
			cfg.GetInt("app.services.post_service.port"),
		), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	postLoaderConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d",
			cfg.GetString("app.services.post_loader.host"),
			cfg.GetInt("app.services.post_loader.port"),
		), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	conns := map[string]interface{}{
		"post_service":        pbp.NewPostServiceClient(postConn),
		"post_loader_service": pbl.NewPostLoaderServiceClient(postLoaderConn),
	}

	return &services{conns: conns}
}

func (s *services) PostService() pbp.PostServiceClient {
	return s.conns["post_service"].(pbp.PostServiceClient)
}

func (s *services) PostLoaderService() pbl.PostLoaderServiceClient {
	return s.conns["post_loader_service"].(pbl.PostLoaderServiceClient)
}
