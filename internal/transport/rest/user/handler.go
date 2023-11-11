package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/service"
	"github.com/moaton/web-api/internal/token"
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
	userSerivce service.UserService
	cache       *cache.Cache
	token       token.Token
}

func NewHandler(userSerivce service.UserService, cache *cache.Cache, token token.Token) Handler {
	return &handler{
		userSerivce: userSerivce,
		cache:       cache,
		token:       token,
	}
}

func (h *handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.token.ExtractIDFromToken(request.RefreshToken)
	if err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.userSerivce.GetUserById(ctx, id)
	if err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}

	accessToken, err := h.token.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.token.CreateRefreshToken(user.ID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.cache.Set(user.ID, refreshToken, time.Now().Add(time.Hour*24).Unix()); err != nil {
		logger.Errorf("Refresh cache.Set err %v", err)
	}

	utils.ResponseOk(w, response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
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

	user, err := h.userSerivce.Auth(ctx, request.Email, request.Password)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, err := h.token.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.token.CreateRefreshToken(user.ID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.cache.Set(user.ID, refreshToken, time.Now().Add(time.Hour*24*7).Unix()); err != nil {
		logger.Errorf("Refresh cache.Set err %v", err)
	}

	utils.ResponseOk(w, response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		logger.Errorf("CreateUser decoder.Decode %v", err)
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.userSerivce.CreateUser(ctx, user)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := h.token.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.token.CreateRefreshToken(user.ID)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
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
