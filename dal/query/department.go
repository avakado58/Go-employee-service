package query

const GetDepartmentByEmployeeId =`
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

const SaveDepartment =`

`

const GetEmployeesDepartmentByEmployeeIds = `

`