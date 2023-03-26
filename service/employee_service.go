package service

import (
	"EnployeeService/dal"
	"EnployeeService/models"
)

type EmployeeService interface {
	SaveEmployee(employee *models.Employee) (id int, err error)
	GetFullEmployee(id int) (employees []models.Employee, err error)
	GetEmployee(id int) (employees []models.Employee, err error)
	UpdateEmployee(employee *models.Employee) (err error)
	DeleteEmployee(id int) (deletedId int, err error)
	GetEmployeesByDepartmentId(id int) (employees []models.Employee, err error)
}

type EmployeeServices struct {
	EmployeeRepository dal.Repository
}

func NewEmployeeService(er dal.Repository) *EmployeeServices {
	return &EmployeeServices{
		EmployeeRepository: er,
	}
}
