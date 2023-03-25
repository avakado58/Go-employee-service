package dal

import (
	dal "EnployeeService/dal/models"
	"EnployeeService/dal/query"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (e *EmployeeRepository)GetDbEmployee(id int) []dal.Employee{
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

	employees := make([]dal.Employee, 0, 1)

	for rows.Next(){
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

func (e *EmployeeRepository) SetDbEmployee(employee *dal.Employee) (id int) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.QueryRow(
		query.SaveEmployee,
		employee.Id,
		employee.Email,
		employee.Age,
		employee.Sex,
		employee.Name,
		employee.Surname,
		employee.Patronymic,
		employee.Birthday,
		employee.PhoneNumber).Scan(&id)
	fmt.Println(id)
	return
}