package revenue

import (
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/moaton/web-api/internal/adapters/api"
	"github.com/moaton/web-api/internal/services"
)

type handler struct {
	revenueService services.RevenueService
}

func NewHandler(service services.RevenueService) api.Handler {
	return &handler{
		revenueService: service,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/revenue", h.GetRevenues).Methods("GET")
	router.HandleFunc("/revenue/{id}", h.GetRevenueById).Methods("GET")
	router.HandleFunc("/revenue/{id}", h.CreateRevenue).Methods("POST")
	router.HandleFunc("/revenue/{id}", h.UpdateRevenue).Methods("PUT")
	router.HandleFunc("/revenue/{id}", h.DeleteRevenue).Methods("DELETE")
}

func (h *handler) GetRevenues(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetRevenueById(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateRevenue(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteRevenue(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateRevenue(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
