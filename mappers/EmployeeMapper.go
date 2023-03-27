package mappers

import (
	"EnployeeService/dal/models"
	"EnployeeService/models"
)

func EmployeeMapToDomain(e dal.Employee) *models.Employee {
	return &models.Employee{
		Id:          e.Id,
		Name:        e.Name,
		Surname:     e.Surname,
		Patronymic:  e.Patronymic,
		Age:         e.Age,
		Sex:         e.Sex,
		PhoneNumber: e.PhoneNumber,
		Email:       e.Email,
		Birthday:    e.Birthday,
	}
}

func EmployeeFullMapToDomain(e dal.Employee) *models.EmployeeFull {
	return &models.EmployeeFull{
		Id:          e.Id,
		Name:        e.Name,
		Surname:     e.Surname,
		Patronymic:  e.Patronymic,
		Age:         e.Age,
		Sex:         e.Sex,
		PhoneNumber: e.PhoneNumber,
		Email:       e.Email,
		Birthday:    e.Birthday,
	}
}

func EmployeeMapToDB(e *models.Employee) *dal.Employee {
	return &dal.Employee{
		Id:          e.Id,
		Name:        e.Name,
		Surname:     e.Surname,
		Patronymic:  e.Patronymic,
		Age:         e.Age,
		Sex:         e.Sex,
		PhoneNumber: e.PhoneNumber,
		Birthday:    e.Birthday,
		Email:       e.Email,
	}
}
