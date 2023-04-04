package query

const GetEmployee = `
select
    e.id as Id,
    e.email as Email,
    e.age as Age,
    e.sex as Sex,
    e.name as Name,
    e.surname as Surname,
    e.patronymic as Patronymic,
    e.birthday as Birthday,
    e.phone_number as PhoneNumber
from employees e
where e.id = $1
`

const GetEmployeeByDepartmentId = `
select
    e.id as Id,
    e.email as Email,
    e.age as Age,
    e.sex as Sex,
    e.name as Name,
    e.surname as Surname,
    e.patronymic as Patronymic,
    e.birthday as Birthday,
    e.phone_number as PhoneNumber
from employees_department ed
	join employees e on e.id = ed.employee_id
where ed.department_id = $1
`

const SaveEmployee = `
insert into employees (name, email, sex, age, surname, patronymic, birthday, phone_number) 
values ($1, $2, $3, $4, $5, $6, $7, $8) 
returning id
`

const UpdateEmployee = `
update employees
set
    email = $1,
	age = $2,
	sex = $3,
	name = $4,
	surname = $5,
	patronymic = $6,
	birthday = $7,
	phone_number = $8
where id = $9
`

const DeleteEmployee = `
delete 
from employees
where id = $1
returning id as deletedId
`
