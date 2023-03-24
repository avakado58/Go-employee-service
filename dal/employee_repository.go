package dal

import (
	"EnployeeService/dal/query"
	"EnployeeService/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (e *EmployeeRepository)GetDbEmployee(id int) *[]models.Employee{
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

	employees := make([]models.Employee, 0, 1)

	for rows.Next(){
		emp := models.Employee{}
		err := rows.Scan(&emp.Id, &emp.Name, &emp.SecondName, &emp.Sex, &emp.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employees = append(employees, emp)
	}

	return &employees
}

func (e *EmployeeRepository) SetDbEmployee(employee *models.Employee) (id int) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.QueryRow(query.SaveEmployee, employee.Name, employee.SecondName, employee.Sex, employee.Age).Scan(&id)
	fmt.Println(id)
	return
}