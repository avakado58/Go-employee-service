package Dal

import (
	"EnployeeService/Models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type EmployeeRepository struct {
	ConnectionString string
}

func NewEmployeeRepository(connectionString string) *EmployeeRepository {
	return &EmployeeRepository{
		ConnectionString: connectionString,
	}
}

func (e *EmployeeRepository)GetDbEmployee(id int) *[]Models.Employee{
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	query := "select e.id as Id, e.name as Name, e.second_name as SecondName, e.sex as Sex, e.age as Age from employees e where id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	employees := make([]Models.Employee, 0, 1)

	for rows.Next(){
		emp := Models.Employee{}
		err := rows.Scan(&emp.Id, &emp.Name, &emp.SecondName, &emp.Sex, &emp.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employees = append(employees, emp)
	}

	return &employees
}

func (e *EmployeeRepository) SetDbEmployee(employee *Models.Employee) (id int) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	query := "insert into employees (name, second_name, sex, age) values ($1, $2, $3, $4) returning id"
	db.QueryRow(query, employee.Name, employee.SecondName, employee.Sex, employee.Age).Scan(&id)
	fmt.Println(id)
	return
}