package models

import (
	"rpc/sharedLib/encoder"
	sl "rpc/sharedLib/models"
	"rpc/sharedLib/std"
)

type EmployeeServer struct{}

func (e *EmployeeServer) AddEmployee(args *[]byte, reply *sl.ApiResponseEmployee) error {
	emp := encoder.UnMarshallEmployee(*args)
	output, err := conn.Query("insert into data.Employee (id, name, secondName, surname, hireDate, position, companyId) values (?, ?, ?, ?, ?, ?, ?)",
		emp.Message.Id, emp.Message.Name, emp.Message.SecondName, emp.Message.Surname, emp.Message.HireDate, emp.Message.Position, emp.Message.CompanyId)
	*reply = emp
	std.ApiLog(err, output, emp, args)
	return nil
}

func (e *EmployeeServer) DeleteEmployee(args *[]byte, reply *sl.ApiResponseEmployee) error {
	emp := encoder.UnMarshallEmployee(*args)
	output, err := conn.Query(" delete from data.Employee where id = ?", emp.Message.Id)
	*reply = emp
	std.ApiLog(err, output, emp, args)
	return nil
}

func (e *EmployeeServer) GetEmployeeById(args *[]byte, reply *sl.ApiResponseEmployee) error {
	emp := encoder.UnMarshallEmployee(*args)
	rows, err := conn.Query("select *from data.Employee where id = ? ", emp.Message.Id)
	std.PrintErr(err)

	for rows.Next() {
		err := rows.Scan(&emp.Message.Id, &emp.Message.Name, &emp.Message.SecondName, &emp.Message.Surname,
			&emp.Message.Position, &emp.Message.HireDate, emp.Message.CompanyId)
		std.PrintErr(err)
	}

	*reply = emp
	std.ApiLog(err, rows, emp, args)
	return nil
}

func (e *EmployeeServer) UpdateEmployeeWithForm(args *[]byte, reply *sl.ApiResponseEmployee) error {
	emp := encoder.UnMarshallEmployee(*args)
	output, err := conn.Query("update data.Employee set name = ?, secondName = ?, surname = ?, position = ?, hireDate = ?, companyId = ? where id = ?",
		emp.Message.Name, emp.Message.SecondName, emp.Message.Surname, emp.Message.Position, emp.Message.HireDate, emp.Message.CompanyId, emp.Message.Id)
	std.PrintErr(err)

	*reply = emp
	std.ApiLog(err, output, emp, args)
	return nil
}

func (e *EmployeeServer) UpdateEmployee(args *[]byte, reply *sl.ApiResponseEmployee) error {
	emp := encoder.UnMarshallEmployee(*args)
	output, err := conn.Query("update data.Employee set name = ?, secondName = ?, surname = ?, position = ?, hireDate = ?, companyId = ? where id = ?",
		emp.Message.Name, emp.Message.SecondName, emp.Message.Surname, emp.Message.Position, emp.Message.HireDate, emp.Message.CompanyId, emp.Message.Id)

	*reply = emp
	std.ApiLog(err, output, emp, args)
	return nil
}
