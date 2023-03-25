package handlers

import (
	"EnployeeService/models"
	"EnployeeService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	es service.EmployeeService
}

func NewEmployeeHandler(es service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		es: es,
	}
}

func (eh *EmployeeHandler)GetFullEmployee(w http.ResponseWriter, r *http.Request)  {
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

func (eh *EmployeeHandler)SaveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	emp := models.Employee{}
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

func (eh *EmployeeHandler)GetEmployee(w http.ResponseWriter, r *http.Request)  {

}

func (eh *EmployeeHandler)UpdateEmployee(w http.ResponseWriter, r *http.Request)  {

}

func (eh *EmployeeHandler)DeleteEmployee(w http.ResponseWriter, r *http.Request)  {

}

func (eh *EmployeeHandler)GetEmployeeByDepartmentId(w http.ResponseWriter, r *http.Request)  {

}