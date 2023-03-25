package dal

import (
	dal "EnployeeService/dal/models"
)

type Repository interface {
	GetDbEmployee(id int) []dal.Employee
	SetDbEmployee(employee *dal.Employee) (id int)
}

type EmployeeRepository struct {
	ConnectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
	}
}