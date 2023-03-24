package Service

import (
	"EnployeeService/Interface"
	"EnployeeService/Models"
)

type EmployeeService struct {
	EmployeeRepository Interface.IEmployeeRepository
}

func NewEmployeeService( er Interface.IEmployeeRepository) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: er,
	}
}

func (e *EmployeeService)SetEmployee(employee *Models.Employee) (id int)  {
	return e.EmployeeRepository.SetDbEmployee(employee)
}

func (e *EmployeeService)GetEmployee(id int) *[]Models.Employee  {
	return e.EmployeeRepository.GetDbEmployee(id)
}