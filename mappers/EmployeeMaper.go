package mappers

import (
	"EnployeeService/dal/models"
	"EnployeeService/models"
)

func MapToDomain(e dal.Employee) *models.Employee{
	return &models.Employee{
		Id:          e.Id,
		Name:        e.Name,
		Surname:  	 e.Surname,
		Patronymic:  e.Patronymic,
		Age:         e.Age,
		Sex:         e.Sex,
		PhoneNumber: e.PhoneNumber,
		Email:     	 e.Email,
		Birthday: 	 e.Birthday,
	}
}

func MapToDB(e *models.Employee) *dal.Employee  {
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
