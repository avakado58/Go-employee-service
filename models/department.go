package models

type Department struct {
	id                                  int
	name                                string
	description                         string
	HeadOfDepartment     				*Employee
	IsClose                            	bool
}
