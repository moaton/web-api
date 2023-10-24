package user

import (
	"encoding/json"
	"net/http"

	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/pkg/logger"
	"github.com/moaton/web-api/pkg/utils"
)

type Handler interface {
	Auth(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userSerivce services.UserService
}

func NewHandler(userSerivce services.UserService) Handler {
	return &handler{
		userSerivce: userSerivce,
	}
}

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		ID int64 `json:"id"`
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.Errorf("CreateUser decoder.Decode %v", err)
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.userSerivce.CreateUser(ctx, user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		ID: id,
	})
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := models.User{}
	h.userSerivce.UpdateUser(ctx, user)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.userSerivce.DeleteUser(ctx, "email@")
	w.WriteHeader(http.StatusOK)
}
