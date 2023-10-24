package revenue

import (
	"net/http"

	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/services"
)

type Handler interface {
	GetRevenues(w http.ResponseWriter, r *http.Request)
	GetRevenueById(w http.ResponseWriter, r *http.Request)
	CreateRevenue(w http.ResponseWriter, r *http.Request)
	DeleteRevenue(w http.ResponseWriter, r *http.Request)
	UpdateRevenue(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	revenueService services.RevenueService
}

func NewHandler(revenueService services.RevenueService) Handler {
	return &handler{
		revenueService: revenueService,
	}
}

func (h *handler) GetRevenues(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = h.revenueService.GetRevenues(ctx, 10, 0)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetRevenueById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = h.revenueService.GetRevenueById(ctx, 0)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	revenue := models.Revenue{}
	h.revenueService.CreateRevenue(ctx, revenue)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.revenueService.DeleteRevenue(ctx, 0)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	revenue := models.Revenue{}
	h.revenueService.UpdateRevenue(ctx, revenue)
	w.WriteHeader(http.StatusOK)
}
