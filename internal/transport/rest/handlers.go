package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/internal/transport/rest/revenue"
	"github.com/moaton/web-api/internal/transport/rest/user"
	"github.com/moaton/web-api/pkg/cache"
)

type Handler interface {
	ListenAndServe()
}

type handler struct {
	userHandler    user.Handler
	revenueHandler revenue.Handler
	middleware     services.MiddleWare
}

func NewHandler(service *services.Service, cache *cache.Cache, middleware services.MiddleWare) *handler {
	return &handler{
		userHandler:    user.NewHandler(service.UserService, cache, middleware),
		revenueHandler: revenue.NewHandler(service.RevenueService, cache, middleware),
	}
}

func (h *handler) ListenAndServe() {
	router := mux.NewRouter()
	router.HandleFunc("/refresh", h.userHandler.Refresh).Methods("POST")

	router.HandleFunc("/user/auth", h.userHandler.Auth).Methods("POST")
	router.HandleFunc("/user", h.userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", h.userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", h.userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/revenue", h.revenueHandler.GetRevenues).Methods("GET")
	router.HandleFunc("/revenue/{id}", h.revenueHandler.GetRevenueById).Methods("GET")
	router.HandleFunc("/revenue", h.revenueHandler.CreateRevenue).Methods("POST")
	router.HandleFunc("/revenue/{id}", h.revenueHandler.UpdateRevenue).Methods("PUT")
	router.HandleFunc("/revenue/{id}", h.revenueHandler.DeleteRevenue).Methods("DELETE")

	if err := http.ListenAndServe(":3030", router); err != nil {
		log.Println("ListenAndServe err ", err)
	}
}
