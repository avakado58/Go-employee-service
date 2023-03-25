package models

import (
	"time"
)

type Employee struct {
	Id              int
	Name            string
	Surname     	string
	Patronymic      string
	Age             int
	Sex             string
	PhoneNumber    	string
	Birthday        time.Time
	Email           string
	Departments		[]Department
}