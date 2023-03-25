package service

import (
	"EnployeeService/mappers"
	"EnployeeService/models"
)

func (e *EmployeeServices)SetEmployee(employee *models.Employee) (id int)  {
	return e.EmployeeRepository.SaveEmployee(mappers.EmployeeMapToDB(employee))
}

func (e *EmployeeServices)GetEmployee(id int) (employees []models.Employee)  {
	dbEmployees := e.EmployeeRepository.GetEmployee(id)

	for _, dbEmp := range dbEmployees{
		emp := mappers.EmployeeMapToDomain(dbEmp)
		dbDeps := e.EmployeeRepository.GetDepartment(dbEmp.Id)
		departments := make([]models.Department, 0, 1)

		for _, dbDep := range dbDeps{
			dep := mappers.DepartmentMapToDomain(dbDep)
			departments = append(departments, *dep)
		}

		emp.Departments = departments
		employees = append(employees, *emp)
	}

	return
}

func (e *EmployeeServices)fillDepartment(emp *models.Employee, id int) (employees []models.Employee) {
	return
}

