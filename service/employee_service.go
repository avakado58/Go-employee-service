package service

import (
	"EnployeeService/dal"
	"EnployeeService/models"
)

type EmployeeService interface {
	SetEmployee(employee *models.Employee) (id int)
	GetEmployee(id int) *[]models.Employee
}

type EmployeeServices struct {
	EmployeeRepository dal.Repository
}

func NewEmployeeService( er dal.Repository) *EmployeeServices {
	return &EmployeeServices{
		EmployeeRepository: er,
	}
}