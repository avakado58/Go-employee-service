package service

import (
	"EnployeeService/mappers"
	"EnployeeService/models"
)

func (e *EmployeeServices)SetEmployee(employee *models.Employee) (id int)  {
	return e.EmployeeRepository.SetDbEmployee(mappers.MapToDB(employee))
}

func (e *EmployeeServices)GetEmployee(id int) []models.Employee  {
	dbEmployees := e.EmployeeRepository.GetDbEmployee(id)
	employees := make([]models.Employee, 0, 1)

	for _, dbEmp := range dbEmployees{
		emp := mappers.MapToDomain(dbEmp)
		employees = append(employees, *emp)
	}

	return employees
}

