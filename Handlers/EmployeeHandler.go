package Handlers

import (
	"EnployeeService/Interface"
	"EnployeeService/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	es Interface.IEmployeeService
}

func NewEmployeeHandler(es Interface.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		es: es,
	}
}

func (eh *EmployeeHandler)GetEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	vars := mux.Vars(r)
	idVar, ok := vars["id"]
	if !ok{
		w.WriteHeader(403)
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	emp := eh.es.GetEmployee(id)
	err = json.NewEncoder(w).Encode(emp)
	if err !=nil{
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}

func (eh *EmployeeHandler)SetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	emp := Models.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil{
		w.WriteHeader(403)
		return
	}

	res := eh.es.SetEmployee(&emp)
	err := json.NewEncoder(w).Encode(res)
	if err !=nil{
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}