package dal

type Department struct {
	Id                                  int
	Name                                string
	Description                         string
	HeadOfDepartmentEmployeeId      	int
	IsClose                            	bool
}