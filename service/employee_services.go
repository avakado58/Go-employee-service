package service

import (
	dal "EnployeeService/dal/models"
	"EnployeeService/mappers"
	"EnployeeService/models"
	"fmt"
)

func (e *EmployeeServices) SaveEmployee(employee *models.Employee) (id int, err error) {
	id = e.EmployeeRepository.SaveEmployee(mappers.EmployeeMapToDB(employee))
	if err != nil {
		fmt.Println(err)
		return
	}

	employee.Id = id
	return id, e.saveEmployeesDepartment(employee)
}

func (e *EmployeeServices) GetFullEmployee(id int) (employees []models.EmployeeFull, err error) {
	dbEmployees, err := e.EmployeeRepository.GetEmployee(id)

	for _, dbEmp := range dbEmployees {
		emp := mappers.EmployeeFullMapToDomain(dbEmp)
		dbDeps := e.DepartmentRepository.GetDepartment(dbEmp.Id)
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

func (e *EmployeeServices) GetEmployee(id int) (employees *models.Employee, err error) {
	dbEmp, err := e.EmployeeRepository.GetEmployee(id)
	if len(dbEmp) == 0 {
		return nil, err
	}

	employees = mappers.EmployeeMapToDomain(dbEmp[0])
	employees.Departments = e.DepartmentRepository.GetDepartmentIdByEmployeeId(employees.Id)

	return
}

func (e *EmployeeServices) UpdateEmployee(employee *models.Employee) (err error) {
	dbEmp, err := e.EmployeeRepository.GetEmployee(employee.Id)
	if err != nil || len(dbEmp) == 0 {
		fmt.Println(err)
		return err
	}

	err = e.EmployeeRepository.UpdateEmployee(mappers.EmployeeMapToDB(employee))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = e.DepartmentRepository.DeleteEmployeesDepartment(dbEmp[0].Id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return e.saveEmployeesDepartment(employee)
}

func (e *EmployeeServices) DeleteEmployee(id int) (deletedId int, err error) {
	err = e.DepartmentRepository.DeleteEmployeesDepartment(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	deletedId, err = e.EmployeeRepository.DeleteEmployee(id)
	return
}

func (e *EmployeeServices) GetEmployeesByDepartmentId(id int) (employees []models.Employee, err error) {
	dbEmployees, err := e.EmployeeRepository.GetEmployeesByDepartmentId(id)
	if len(dbEmployees) == 0 {
		return []models.Employee{}, err
	}

	for _, dbEmp := range dbEmployees {
		emp := mappers.EmployeeMapToDomain(dbEmp)
		emp.Departments = e.DepartmentRepository.GetDepartmentIdByEmployeeId(emp.Id)
		employees = append(employees, *emp)
	}

	return
}

func (e *EmployeeServices) saveEmployeesDepartment(employee *models.Employee) (err error) {
	var empDeps []dal.EmployeesDepartment

	for _, depId := range employee.Departments {
		empDeps = append(empDeps, dal.EmployeesDepartment{
			EmployeeId:   employee.Id,
			DepartmentId: depId,
		})
	}

	_, err = e.DepartmentRepository.SaveEmployeesDepartment(empDeps)
	if err != nil {
		fmt.Println(err)
	}

	return
}
