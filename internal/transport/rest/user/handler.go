package user

import (
	"encoding/json"
	"net/http"

	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/pkg/cache"
	"github.com/moaton/web-api/pkg/logger"
	"github.com/moaton/web-api/pkg/utils"
)

type Handler interface {
	Refresh(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	userSerivce services.UserService
	cache       *cache.Cache
	middleware  services.MiddleWare
}

func NewHandler(userSerivce services.UserService, cache *cache.Cache, middleware services.MiddleWare) Handler {
	return &handler{
		userSerivce: userSerivce,
		cache:       cache,
		middleware:  middleware,
	}
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	type response struct {
		Access  string `json:"access_token"`
		Refresh string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.userSerivce.Refresh(ctx, request.RefreshToken)
	if err != nil {
		logger.Errorf("Refresh err %v", err)
		utils.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Access:  tokens["access_token"],
		Refresh: tokens["refresh_token"],
	})
}

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	if request.Email == "" || request.Password == "" {
		utils.ResponseError(w, http.StatusBadRequest, "The email or password is empty")
		return
	}

	user, err := h.userSerivce.GetUserByEmail(ctx, request.Email, request.Password)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseOk(w, response{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	})
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Refresh string `json:"refresh_token"`
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.Errorf("CreateUser decoder.Decode %v", err)
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.userSerivce.CreateUser(ctx, user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	cookie := http.Cookie{
		Name:     "Bearer",
		Value:    tokens["access_token"],
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)

	utils.ResponseOk(w, response{
		Refresh: tokens["refresh_token"],
	})
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type response struct {
		Message string `json:"message"`
	}
	user := models.User{}
	err := h.userSerivce.UpdateUser(ctx, user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Message: "success",
	})
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type response struct {
		Message string `json:"message"`
	}
	var request struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	if request.Email == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Email is empty")
		return
	}
	err = h.userSerivce.DeleteUser(ctx, request.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Message: "success",
	})
}
