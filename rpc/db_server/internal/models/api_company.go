package models

import (
	"fmt"
	"os"
	"rpc/sharedLib/encoder"
	sl "rpc/sharedLib/models"
	"rpc/sharedLib/std"
)

type CompanyServer struct{}

func (s *CompanyServer) DeleteCompany(args *[]byte, reply *sl.ApiResponseCompany) error {
	cmp := encoder.UnMarshallCompany(*args)
	output, err := conn.Query("delete from data.Company where id = ?", cmp.Message.Id)
	*reply = cmp
	std.ApiLog(err, output, cmp, args)
	return nil
}

func (s *CompanyServer) GetCompanyById(args *[]byte, reply *sl.ApiResponseCompany) error {
	cmp := encoder.UnMarshallCompany(*args)
	rows, err := conn.Query("select *from data.Company where id = ?", cmp.Message.Id)
	std.PrintErr(err)

	for rows.Next() {
		parse := rows.Scan(&cmp.Message.Id, &cmp.Message.Name, &cmp.Message.LegalForm)
		fmt.Println(parse)
	}

	*reply = cmp
	std.ApiLog(err, rows, cmp, reply)
	return nil
}

func (s *CompanyServer) GetEmployeesByCompanyId(args *[]byte, reply *sl.MyNotFavouriteUniqueStruct) error {
	cmp := encoder.UnMarshallCompany(*args)
	fmt.Fprintf(os.Stderr, "%v\n", cmp.Message.Id)
	rows, err := conn.Query("select *from data.Employee where companyId = ?", cmp.Message.Id)
	std.PrintErr(err)

	var emp sl.Employee
	arr := make([]sl.Employee, 1)
	for rows.Next() {
		err := rows.Scan(&emp.Id, &emp.Name, &emp.SecondName, &emp.Surname, &emp.HireDate, &emp.Position, &emp.CompanyId)
		arr = append(arr, emp)
		std.PrintErr(err)
	}

	reply.Payload.Emp = arr
	std.ApiLog(err, rows, cmp, args)
	fmt.Println(arr)
	return nil
}

func (s *CompanyServer) PlaceCompany(args *[]byte, reply *sl.ApiResponseCompany) error {
	cmp := encoder.UnMarshallCompany(*args)
	rows, err := conn.Query("insert into data.Company (id, name, legalForm) values (?, ?, ?)",
		cmp.Message.Id, cmp.Message.Name, cmp.Message.LegalForm)

	*reply = cmp
	std.ApiLog(err, rows, cmp, args)
	return nil
}

func (s *CompanyServer) UpdateCompanyWithForm(args *[]byte, reply *sl.ApiResponseCompany) error {
	cmp := encoder.UnMarshallCompany(*args)
	output, err := conn.Query("update data.Company set name = ?, legalForm = ? where id = ?",
		cmp.Message.Name, cmp.Message.LegalForm, cmp.Message.Id)

	*reply = cmp
	std.ApiLog(err, output, cmp, args)
	return nil
}

func (s *CompanyServer) UpdateCompany(args *[]byte, reply *sl.ApiResponseCompany) error {
	cmp := encoder.UnMarshallCompany(*args)
	output, err := conn.Query("update data.Company set name = ?, legalForm = ? where id = ?",
		cmp.Message.Name, cmp.Message.LegalForm, cmp.Message.Id)
	*reply = cmp
	std.ApiLog(err, output, cmp, args)
	return nil
}
