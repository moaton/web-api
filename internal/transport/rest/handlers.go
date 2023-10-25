package adapters

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moaton/web-api/internal/services"
	"github.com/moaton/web-api/internal/transport/rest/revenue"
	"github.com/moaton/web-api/internal/transport/rest/user"
)

type Handler interface {
	Register(router *mux.Router)
}

type handler struct {
	userHandler    user.Handler
	revenueHandler revenue.Handler
}

func ListenAndServe(service *services.Service) {
	router := mux.NewRouter()

	handler := &handler{
		userHandler:    user.NewHandler(service.UserService),
		revenueHandler: revenue.NewHandler(service.RevenueService),
	}

	router.HandleFunc("/user/auth", handler.userHandler.Auth).Methods("POST")
	router.HandleFunc("/user", handler.userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", handler.userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.userHandler.DeleteUser).Methods("DELETE")

	router.HandleFunc("/revenue", handler.revenueHandler.GetRevenues).Methods("GET")
	router.HandleFunc("/revenue/{id}", handler.revenueHandler.GetRevenueById).Methods("GET")
	router.HandleFunc("/revenue/{id}", handler.revenueHandler.CreateRevenue).Methods("POST")
	router.HandleFunc("/revenue/{id}", handler.revenueHandler.UpdateRevenue).Methods("PUT")
	router.HandleFunc("/revenue/{id}", handler.revenueHandler.DeleteRevenue).Methods("DELETE")

	if err := http.ListenAndServe(":3030", router); err != nil {
		log.Println("ListenAndServe err ", err)
	}
}
