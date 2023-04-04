package mappers

import (
	dal "EnployeeService/dal/models"
	"EnployeeService/models"
)

func DepartmentMapToDomain(d dal.Department) *models.Department  {
	return &models.Department{
		Id:               d.Id,
		Name:             d.Name,
		Description:      d.Description,
		HeadOfDepartment: nil,
		IsClose:          d.IsClose,
	}
}