package dal

import (
	"EnployeeService/dal/models"
	"EnployeeService/dal/query"
	"database/sql"
	"fmt"
)

func (e *DepartmentRepository) GetDepartment(employeeId int) (department []dal.Department) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, err := db.Query(query.GetDepartmentByEmployeeId, employeeId)

	for rows.Next() {
		dep := dal.Department{}
		err := rows.Scan(&dep.Id, &dep.Name, &dep.Description, &dep.HeadOfDepartmentEmployeeId, &dep.IsClose)
		if err != nil {
			fmt.Println(err)
			continue
		}

		department = append(department, dep)
	}

	defer rows.Close()

	return
}

func (e *DepartmentRepository) GetDepartmentIdByEmployeeId(employeeId int) (depIds []int) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, err := db.Query(query.GetDepartmentIdsByEmployeeIds, employeeId)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		depIds = append(depIds, id)
	}

	return
}

func (e *DepartmentRepository) GetEmployeesDepartment(employeeIds int) (employeesDepartment []dal.EmployeesDepartment) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, err := db.Query(query.GetEmployeesDepartmentByEmployeeIds, employeeIds)

	for rows.Next() {
		dep := dal.EmployeesDepartment{}
		err = rows.Scan(&dep.Id, &dep.EmployeeId, &dep.DepartmentId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		employeesDepartment = append(employeesDepartment, dep)
	}

	defer rows.Close()

	return
}

func (e *DepartmentRepository) DeleteEmployeesDepartment(employeeIds int) (err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Query(query.DeleteEmployeesDepartmentByEmployeeIds, employeeIds)

	return
}

func (e *DepartmentRepository) SaveEmployeesDepartment(ed []dal.EmployeesDepartment) (count int, err error) {
	db, err := sql.Open("postgres", e.ConnectionString)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var i int
	for _, item := range ed {
		_, err = db.Query(query.SaveEmployeesDepartment, item.EmployeeId, item.DepartmentId)
		if err != nil {
			fmt.Println(err)
			return i, err
		}

		i++
	}

	return i, err
}
