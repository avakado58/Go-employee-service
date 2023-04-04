package dal

import (
	dal "EnployeeService/dal/models"
	"EnployeeService/dal/query"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (e *EmployeeRepository) GetEmployee(id int) (employees []dal.Employee, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rows, err := db.Query(query.GetEmployee, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	employees = e.scanEmployees(rows, employees)
	return
}

func (e *EmployeeRepository) GetEmployeesByDepartmentId(id int) (employees []dal.Employee, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	rows, err := db.Query(query.GetEmployeeByDepartmentId, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	employees = e.scanEmployees(rows, employees)
	return
}

func (e *EmployeeRepository) SaveEmployee(employee *dal.Employee) (id int) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.QueryRow(
		query.SaveEmployee,
		employee.Name,
		employee.Email,
		employee.Sex,
		employee.Age,
		employee.Surname,
		employee.Patronymic,
		employee.Birthday,
		employee.PhoneNumber).Scan(&id)

	return
}

func (e *EmployeeRepository) UpdateEmployee(employee *dal.Employee) (err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer db.Close()

	db.QueryRow(
		query.UpdateEmployee,
		employee.Email,
		employee.Age,
		employee.Sex,
		employee.Name,
		employee.Surname,
		employee.Patronymic,
		employee.Birthday,
		employee.PhoneNumber,
		employee.Id)

	return nil
}

func (e *EmployeeRepository) DeleteEmployee(id int) (deletedId int, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer db.Close()

	db.QueryRow(query.DeleteEmployee, id).Scan(&deletedId)

	return
}

func (e *EmployeeRepository) scanEmployees(rows *sql.Rows, employees []dal.Employee) []dal.Employee {
	for rows.Next() {
		emp := dal.Employee{}
		err := rows.Scan(&emp.Id, &emp.Email, &emp.Age, &emp.Sex, &emp.Name, &emp.Surname, &emp.Patronymic, &emp.Birthday, &emp.PhoneNumber)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employees = append(employees, emp)
	}

	return employees
}
