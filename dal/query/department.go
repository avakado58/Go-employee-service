package query

const GetDepartmentByEmployeeId = `
select
        d.id as Id,
        d.name as Name,
        d.description as Description,
        d.head_of_department_employee_id as HeadOfDepartmentEmployeeId,
        d.is_close as IsClose
from departments d
join employees_department ed on d.id = ed.department_id
where ed.employee_id = $1
`

const SaveDepartment = `

`

const GetDepartmentIdsByEmployeeIds = `
select
       d.id as id
from departments d
    join employees_department ed on d.id = ed.department_id
where ed.employee_id = $1
`

const GetEmployeesDepartmentByEmployeeIds = `
select
       ed.id as Id,
       ed.employee_id as EmployeeId,
       ed.department_id as DepartmentId
from employees_department ed
where ed.employee_id = $1
`

const DeleteEmployeesDepartmentByEmployeeIds = `
delete
from employees_department
where employee_id = $1
`

const SaveEmployeesDepartment = `
insert into employees_department (employee_id, department_id)
values ($1, $2)
`
