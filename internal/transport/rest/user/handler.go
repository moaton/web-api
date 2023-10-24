package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/services"
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
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("CreateUser decoder.Decode")
	}
	h.userSerivce.CreateUser(ctx, user)
	w.WriteHeader(http.StatusOK)
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
