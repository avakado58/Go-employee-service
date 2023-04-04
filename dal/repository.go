package dal

import (
	"EnployeeService/dal/models"
)

type IEmployeeRepository interface {
	GetEmployee(id int) (employees []dal.Employee, err error)
	SaveEmployee(employee *dal.Employee) (id int)
	UpdateEmployee(employee *dal.Employee) (err error)
	DeleteEmployee(id int) (deletedId int, err error)
	GetEmployeesByDepartmentId(id int) (employees []dal.Employee, err error)
}

type IDepartmentRepository interface {
	GetDepartment(employeeId int) (department []dal.Department)
	GetDepartmentIdByEmployeeId(employeeId int) (depIds []int)
	GetEmployeesDepartment(employeeIds int) (employeesDepartment []dal.EmployeesDepartment)
	DeleteEmployeesDepartment(employeeIds int) (err error)
	SaveEmployeesDepartment(ed []dal.EmployeesDepartment) (count int, err error)
}

type EmployeeRepository struct {
	ConnectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
	}
}

type DepartmentRepository struct {
	ConnectionString string
}

func NewDepartmentRepository(connectionString string) *DepartmentRepository {
	return &DepartmentRepository{
		ConnectionString: connectionString,
	}
}
