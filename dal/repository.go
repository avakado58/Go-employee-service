package dal

import (
	"EnployeeService/models"
)

type Repository interface {
	GetDbEmployee(id int) *[]models.Employee
	SetDbEmployee(employee *models.Employee) (id int)
}

type EmployeeRepository struct {
	ConnectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
	}
}