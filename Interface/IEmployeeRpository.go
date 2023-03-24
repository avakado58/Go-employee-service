package Interface

import (
	"EnployeeService/Models"
)

type IEmployeeRepository interface {
	GetDbEmployee(id int) *[]Models.Employee
	SetDbEmployee(employee *Models.Employee) (id int)
}

