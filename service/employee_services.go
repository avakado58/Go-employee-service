package service

import (
	"EnployeeService/models"
)

func (e *EmployeeServices)SetEmployee(employee *models.Employee) (id int)  {
	return e.EmployeeRepository.SetDbEmployee(employee)
}

func (e *EmployeeServices)GetEmployee(id int) *[]models.Employee  {
	return e.EmployeeRepository.GetDbEmployee(id)
}