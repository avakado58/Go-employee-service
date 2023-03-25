package api

import (
	"EnployeeService/api/handlers"
	"context"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func NewServer(handler *handlers.EmployeeHandler, ctx context.Context, port string) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/employee/{id}", handler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee", handler.SetEmployee).Methods(http.MethodPost)

	return &http.Server{
		Addr:              "",
		Handler:           r,
		BaseContext:       func(_ net.Listener) context.Context {
			return ctx
		},
	}
}