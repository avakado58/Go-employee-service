package Interface

import (
	"EnployeeService/Models"
)

type IEmployeeService interface {
	SetEmployee(employee *Models.Employee) (id int)
	GetEmployee(id int) *[]Models.Employee
}
