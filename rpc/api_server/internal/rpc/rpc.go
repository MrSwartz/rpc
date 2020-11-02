package rpc

import (
	"log"
	"net/http"
	"net/rpc"
	"rpc/sharedLib/encoder"
	"rpc/sharedLib/help"
	"rpc/sharedLib/models"
	"rpc/sharedLib/std"
)

func RPCConn() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", ":1234")
	log.Println(err)
	return client
}

var client = RPCConn()

func UpdateEmployeeWithForm(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseEmployee{}, &models.ApiResponseEmployee{}

	args.Message = help.ParseURLArgsEmployee(r)
	data := encoder.MarshallEmployee(args.Message, http.MethodPut, http.StatusOK)

	err := client.Call("EmployeeServer.UpdateEmployee", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseEmployee{}, &models.ApiResponseEmployee{}

	args.Message = help.ParseURLArgsEmployee(r)
	data := encoder.MarshallEmployee(args.Message, http.MethodDelete, http.StatusOK)

	err := client.Call("EmployeeServer.DeleteEmployee", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseEmployee{}, &models.ApiResponseEmployee{}

	args.Message = help.ParseURLArgsEmployee(r)
	data := encoder.MarshallEmployee(args.Message, http.MethodPut, http.StatusOK)

	err := client.Call("EmployeeServer.UpdateEmployee", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseEmployee{}, &models.ApiResponseEmployee{}

	args.Message = help.ParseURLArgsEmployee(r)
	data := encoder.MarshallEmployee(args.Message, http.MethodPost, http.StatusOK)

	err := client.Call("EmployeeServer.AddEmployee", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseEmployee{}, &models.ApiResponseEmployee{}

	args.Message.Id = help.GetId(r)
	data := encoder.MarshallEmployee(args.Message, http.MethodGet, http.StatusOK)

	err := client.Call("EmployeeServer.GetEmployeeById", data, &reply)
	std.PrintLogs(err, r, data, reply)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseCompany{}, &models.ApiResponseCompany{}

	args.Message.Id = help.GetId(r)
	data := encoder.MarshallCompany(args.Message, http.MethodDelete, http.StatusOK)

	err := client.Call("CompanyServer.DeleteCompany", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func UpdateCompanyWithForm(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseCompany{}, &models.ApiResponseCompany{}

	args.Message = help.ParseURLArgsCompany(r)
	data := encoder.MarshallCompany(args.Message, http.MethodPost, http.StatusOK)

	err := client.Call("CompanyServer.UpdateCompanyWithForm", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func GetCompanyById(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseCompany{}, &models.ApiResponseCompany{}

	args.Message.Id = help.GetId(r)
	data := encoder.MarshallCompany(args.Message, http.MethodGet, http.StatusOK)

	err := client.Call("CompanyServer.GetCompanyById", data, &reply)
	std.PrintLogs(err, r, data, reply)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseCompany{}, &models.ApiResponseCompany{}

	args.Message = help.ParseURLArgsCompany(r)
	data := encoder.MarshallCompany(args.Message, http.MethodPost, http.StatusOK)

	err := client.Call("CompanyServer.UpdateCompany", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func PlaceCompany(w http.ResponseWriter, r *http.Request) {
	reply, args := models.ApiResponseCompany{}, &models.ApiResponseCompany{}

	args.Message = help.ParseURLArgsCompany(r)
	data := encoder.MarshallCompany(args.Message, http.MethodPost, http.StatusOK)

	err := client.Call("CompanyServer.PlaceCompany", data, &reply)
	std.PrintLogs(err, r, args, reply)
}

func GetEmployeesByCompanyId(w http.ResponseWriter, r *http.Request) {
	reply, args := models.MyNotFavouriteUniqueStruct{}, &models.ApiResponseCompany{}

	args.Message.Id = help.GetId(r)
	data := encoder.MarshallCompany(args.Message, http.MethodGet, http.StatusOK)

	err := client.Call("CompanyServer.GetEmployeesByCompanyId", data, &reply)
	std.PrintLogs(err, r, args, reply)

}
