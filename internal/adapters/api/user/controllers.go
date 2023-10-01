package revenue

import (
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/moaton/web-api/internal/adapters/api"
	"github.com/moaton/web-api/internal/entities/revenue"
)

type handler struct {
	revenueService revenue.Service
}

func NewHandler(service revenue.Service) api.Handler {
	return &handler{
		revenueService: service,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/user/auth", h.Auth).Methods("POST")
	router.HandleFunc("/user/", h.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
