package main

import (
	"EnployeeService/api/handlers"
	"EnployeeService/config"
	"EnployeeService/dal"
	"EnployeeService/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg := config.ReadCfg()
	log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", cfg.Port, cfg.Env))

	repos := dal.NewEmployeeRepository(cfg.ConnectionString)
	emplService := service.NewEmployeeService(repos)
	emplHandler := handlers.NewEmployeeHandler(emplService)

	r := mux.NewRouter()
	r.HandleFunc("/employee/{id}", emplHandler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee", emplHandler.SetEmployee).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(cfg.Port, nil)
}
