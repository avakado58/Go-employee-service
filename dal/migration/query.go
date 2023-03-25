package migration

const AddEmployeeTable = `
create table employees(
    id              serial primary key,
    name            varchar(100),
	surname     varchar(100),
	patronymic         varchar(100),
	age             int,
	sex             varchar(10),
	phone_number    varchar(20),
	birthday        timestamp,
	email           varchar(100)
);
create index employees_employee_id on employees(id);
`

const AddDepartmentTable = `
create table department(
    id                                  serial primary key,
    name                                varchar(100),
    description                         varchar(100),
    head_of_department_employee_id      int,
    is_close                            bool default false
);
create index departments_department_id on departments(id);
alter table departments
add constraint fk_employees foreign key(head_of_department_employee_id) references employees(id);
`

const AddEmployeesDepartmentTable = `
create table employees_department(
    id                  serial primary key,
    employee_id         int,
    department_id       int
);
create index employees_departments_id on employees_department(id);
alter table employees_department
add constraint fk_employees_department_employee_id foreign key(employee_id) references employees(id),
add constraint fk_employees_department_department_id foreign key(department_id) references departments(id);
`