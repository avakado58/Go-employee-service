package query

const GetEmployee = `
select 
	e.id as Id, 
	e.name as Name, 
	e.second_name as SecondName, 
	e.sex as Sex, 
	e.age as Age 
from employees e 
where id = $1
`

const SaveEmployee = `
insert into employees (name, second_name, sex, age) 
values ($1, $2, $3, $4) 
returning id
`