package dal

type Department struct {
	id                                  int
	name                                string
	description                         string
	head_of_department_employee_id      int
	is_close                            bool
}