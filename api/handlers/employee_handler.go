package handlers

import (
	"EnployeeService/models"
	"EnployeeService/service"
	"encoding/json"
	"fmt"
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

func (eh *EmployeeHandler) GetFullEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	var id int
	id = eh.getParam(r, w, "id", id).(int)
	emp, err := eh.es.GetFullEmployee(id)
	if err != nil {
		w.WriteHeader(404)
	}

	eh.writeResponse(w, emp)
}

func (eh *EmployeeHandler) SaveEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	emp := models.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.WriteHeader(403)
		return
	}

	res, err := eh.es.SaveEmployee(&emp)
	if err != nil {
		w.WriteHeader(500)
	}

	eh.writeResponse(w, res)
}

func (eh *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;encoding=utf-8")
	var id int
	id = eh.getParam(r, w, "id", id).(int)

	emp, err := eh.es.GetEmployee(id)
	if err != nil {
		w.WriteHeader(404)
	}

	eh.writeResponse(w, emp)
}

func (eh *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

}

func (eh *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

}

func (eh *EmployeeHandler) GetEmployeeByDepartmentId(w http.ResponseWriter, r *http.Request) {

}

func (eh *EmployeeHandler) getParam(r *http.Request, w http.ResponseWriter, name string, t interface{}) (param interface{}) {
	vars := mux.Vars(r)
	v, ok := vars[name]
	if !ok {
		w.WriteHeader(403)
		return
	}

	var err error

	switch t.(type) {
	case int:
		param, err = strconv.Atoi(v)
	case float64:
		param, err = strconv.ParseFloat(v, 64)
	default:
		fmt.Println("unknown")
	}

	if err != nil {
		w.WriteHeader(403)
		return
	}

	return
}

func (eh *EmployeeHandler) writeResponse(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
}
