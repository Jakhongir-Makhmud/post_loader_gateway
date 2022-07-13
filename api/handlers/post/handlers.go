package post

import (
	"api-gateway-iman/api/structs"
	pbp "api-gateway-iman/genproto/post_service"
	"api-gateway-iman/pkg/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (h *postHandler) GetPost(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		utils.ReplyToReq(rw, http.StatusBadRequest, pbp.Post{})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ReplyToReq(rw, http.StatusBadRequest, pbp.Post{})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()
	post, err := h.services.PostService().GetPost(ctx, &pbp.PostId{Id: int64(id)})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			utils.ReplyToReq(rw, http.StatusNotFound, structs.NotFoundResponse)
			return
		}
		h.Logger.Error("can not get post from post-service", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
		return
	}

	utils.ReplyToReq(rw, http.StatusOK, post)
}

func (h *postHandler) UpdatePost(rw http.ResponseWriter, r *http.Request) {

	var request structs.UpdatePostReq

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.Logger.Warn("can not unmarshal json to struct", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusBadRequest, pbp.Post{})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()

	post, err := h.services.PostService().UpdatePost(ctx, &pbp.Post{Id: request.Id, Title: request.Title, Body: request.Body})
	if err != nil {
		h.Logger.Error("can not update post", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
		return
	}

	utils.ReplyToReq(rw, http.StatusOK, post)
}

func (h *postHandler) DeletePost(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		utils.ReplyToReq(rw, http.StatusBadRequest, struct{}{})
		return
	}
	id, err := strconv.Atoi(idStr)
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()

	_, err = h.services.PostService().DeletePost(ctx, &pbp.PostId{Id: int64(id)})

	if err != nil {
		h.Logger.Error("can not delete post", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
	}

}

func (h *postHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	var request structs.GetListPostsReq

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.Logger.Warn("can not unmarshal json to struct", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusBadRequest, []pbp.Post{})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(h.cfg.GetInt("app.services.reqTimeout"))*time.Second)
	defer cancel()

	posts, err := h.services.PostService().ListPost(ctx, &pbp.ListOfPosts{Page: request.Page, Limit: request.Limit})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			utils.ReplyToReq(rw, http.StatusNotFound, structs.NotFoundResponse)
			return
		}
		h.Logger.Error("can not get posts from post-service", zap.Error(err))
		utils.ReplyToReq(rw, http.StatusInternalServerError, structs.ErrInternalResponse)
		return
	}

	utils.ReplyToReq(rw, http.StatusOK, posts)
}
