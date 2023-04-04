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
	r.HandleFunc("/employee/{id}/full", handler.GetFullEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", handler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee", handler.SaveEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee", handler.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", handler.DeleteEmployee).Methods(http.MethodDelete)
	r.HandleFunc("/employee/department/{id}", handler.GetEmployeeByDepartmentId).Methods(http.MethodGet)

	return &http.Server{
		Addr:              "",
		Handler:           r,
		BaseContext:       func(_ net.Listener) context.Context {
			return ctx
		},
	}
}