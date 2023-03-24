package main

import (
	"EnployeeService/Config"
	"EnployeeService/Dal"
	"EnployeeService/Handlers"
	"EnployeeService/Service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg := Config.ReadCfg()
	log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", cfg.Port, cfg.Env))

	repos := Dal.NewEmployeeRepository(cfg.ConnectionString)
	emplService := Service.NewEmployeeService(repos)
	emplHandler := Handlers.NewEmployeeHandler(emplService)

	r := mux.NewRouter()
	r.HandleFunc("/employee/{id}", emplHandler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee", emplHandler.SetEmployee).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(cfg.Port, nil)
}
