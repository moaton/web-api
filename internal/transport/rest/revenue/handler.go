package revenue

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/pkg/utils"
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

	var request struct {
		Limit  int64  `json:"limit"`
		Offset int64  `json:"offset"`
		Query  string `json:"query"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	type response struct {
		Revenues []models.Revenue `json:"revenues"`
		Total    int64            `json:"total"`
	}

	revenues, total, err := h.revenueService.GetRevenues(ctx, request.Limit, request.Offset)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Revenues: revenues,
		Total:    total,
	})
}

func (h *handler) GetRevenueById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Revenue models.Revenue `json:"revenue"`
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	revenue, err := h.revenueService.GetRevenueById(ctx, id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseOk(w, response{
		Revenue: revenue,
	})
}

func (h *handler) CreateRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		ID int64 `json:"id"`
	}

	var revenue models.Revenue

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&revenue)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if revenue.Title == "" || revenue.Type == "" {
		utils.ResponseError(w, http.StatusBadRequest, "title or type is empty")
		return
	}

	id, err := h.revenueService.CreateRevenue(ctx, revenue)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		ID: id,
	})
}

func (h *handler) DeleteRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Message string `json:"message"`
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.revenueService.DeleteRevenue(ctx, id)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ResponseOk(w, response{
		Message: "success",
	})
}

func (h *handler) UpdateRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Message string `json:"message"`
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var revenue models.Revenue

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&revenue)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	revenue.ID = id

	err = h.revenueService.UpdateRevenue(ctx, revenue)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseOk(w, response{
		Message: "success",
	})
}
