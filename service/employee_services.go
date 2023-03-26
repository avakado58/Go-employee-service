package service

import (
	"EnployeeService/mappers"
	"EnployeeService/models"
)

func (e *EmployeeServices) SaveEmployee(employee *models.Employee) (id int, err error) {
	return e.EmployeeRepository.SaveEmployee(mappers.EmployeeMapToDB(employee)), nil
}

func (e *EmployeeServices) GetFullEmployee(id int) (employees []models.Employee, err error) {
	dbEmployees := e.EmployeeRepository.GetEmployee(id)

	for _, dbEmp := range dbEmployees {
		emp := mappers.EmployeeMapToDomain(dbEmp)
		dbDeps := e.EmployeeRepository.GetDepartment(dbEmp.Id)
		departments := make([]models.Department, 0, 1)

		for _, dbDep := range dbDeps {
			dep := mappers.DepartmentMapToDomain(dbDep)
			departments = append(departments, *dep)
		}

		emp.Departments = departments
		employees = append(employees, *emp)
	}

	return
}

func (e *EmployeeServices) GetEmployee(id int) (employees []models.Employee, err error) {
	return
}

func (e *EmployeeServices) UpdateEmployee(employee *models.Employee) (err error) {
	return
}

func (e *EmployeeServices) DeleteEmployee(id int) (deletedId int, err error) {
	return
}

func (e *EmployeeServices) GetEmployeesByDepartmentId(id int) (employees []models.Employee, err error) {
	return
}
