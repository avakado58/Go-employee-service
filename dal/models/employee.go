package dal

import "time"

type Employee struct {
	id              int
	name            string
	surname     	string
	patronymic      string
	age             int
	sex             string
	phone_number    string
	birthday        time.Time
	email           string
}