package dal

import (
	dal "EnployeeService/dal/models"
)

type Repository interface {
	GetEmployee(id int) []dal.Employee
	SaveEmployee(employee *dal.Employee) (id int)
	GetDepartment(employeeId int) (department []dal.Department)
	SaveDepartment(department *dal.Department) (id int)
	GetEmployeesDepartment(employeeIds []int) (employeesDepartment []dal.EmployeesDepartment)
}

type EmployeeRepository struct {
	ConnectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
	}
}