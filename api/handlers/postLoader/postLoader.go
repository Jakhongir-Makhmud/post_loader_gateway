package postloader

import (
	"api-gateway-iman/api/structs"
	pbl "api-gateway-iman/genproto/post_loader_service"
	serviceconn "api-gateway-iman/internal/serviceConn"
	"api-gateway-iman/pkg/config"
	"api-gateway-iman/pkg/logger"
	"api-gateway-iman/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Params struct {
	Cfg      config.Config
	Logger   logger.Logger
	Services serviceconn.ServiceController
}

type PostLoaderHandler interface {
	LoadPosts(rw http.ResponseWriter, r *http.Request)
	GetLoadingStatus(rw http.ResponseWriter, r *http.Request)
}

type postLoaderHandler struct {
	cfg      config.Config
	logger   logger.Logger
	services serviceconn.ServiceController
}

func NewPostLoader(params Params) PostLoaderHandler {

	return &postLoaderHandler{
		cfg:      params.Cfg,
		logger:   params.Logger,
		services: params.Services,
	}
}

func (h *postLoaderHandler) LoadPosts(rw http.ResponseWriter, r *http.Request) {
	var request structs.LoadPostsReq

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Warn("can not unmarshal json to struct", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusBadRequest, pbl.LoadingStatus{})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()

	status, err := h.services.PostLoaderService().LoadPosts(ctx, &pbl.LoadPostParam{Pages: request.Pages})
	if err != nil {
		h.logger.Error("failed to load posts into database", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
		return
	}

	utils.ReplyToReq(rw, http.StatusOK, status)
}

func (h *postLoaderHandler) GetLoadingStatus(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		utils.ReplyToReq(rw, http.StatusBadRequest, pbl.LoadingStatus{})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()

	status, err := h.services.PostLoaderService().GetJobStatus(ctx, &pbl.JobId{Id: id})
	if err != nil {
		h.logger.Error("can not get loading status", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
		return
	}

	utils.ReplyToReq(rw, http.StatusOK, status)
}
