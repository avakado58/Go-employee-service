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

const SaveEmployee = `
insert into employees (name, email, sex, age, surname, patronymic, birthday, phone_number) 
values ($1, $2, $3, $4, $5, $6, $7, $8) 
returning id
`