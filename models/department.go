package models

type Department struct {
	Id                                  int
	Name                                string
	Description                         string
	HeadOfDepartment     				*Employee
	IsClose                            	bool
}
