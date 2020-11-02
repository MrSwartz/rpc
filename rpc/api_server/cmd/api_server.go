package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rpc/api_server/internal/config"
	"rpc/api_server/internal/rpc"
)

func main() {
	r := mux.NewRouter()
	cnf := config.New()
																														//Пример запроса в Postman
	r.HandleFunc(cnf.Path+"/company/", rpc.UpdateCompany).Methods(http.MethodPut)                                		//http://localhost:8080/hostname/company/?Id=1&Name="Linus"&LegalForm="lf1"
	r.HandleFunc(cnf.Path+"/company/", rpc.PlaceCompany).Methods(http.MethodPost)                                		//http://localhost:8080/hostname/company/1?Id=1&Name="Linus"&LegalForm="lf1"
	r.HandleFunc(cnf.Path+"/company/{companyId}", rpc.UpdateCompanyWithForm).Methods(http.MethodPost)            		//http://localhost:8080/hostname/company/1?Id=1&Name="Linus"&LegalForm="lf1"
	r.HandleFunc(cnf.Path+"/company/{companyId}", rpc.GetCompanyById).Methods(http.MethodGet)                    		//http://localhost:8080/hostname/company/1?Id=1&Name="Linus"
	r.HandleFunc(cnf.Path+"/company/{companyId}", rpc.DeleteCompany).Methods(http.MethodDelete)                  		//http://localhost:8080/hostname/company/1?Id=1
	r.HandleFunc(cnf.Path+"/company/{companyId}/employees", rpc.GetEmployeesByCompanyId).Methods(http.MethodGet) 		//http://localhost:8080/hostname/company/1/employees?Id=1

	r.HandleFunc(cnf.Path+"/employee", rpc.AddEmployee).Methods(http.MethodPost)                        				//http://localhost:8080/hostname/employee?Id=1&Name="Linus"&SecondName="Torvalds"&Surname="Benedict"&Position=developer&HireDate=2017-03-03&CompanyId=1
	r.HandleFunc(cnf.Path+"/employee", rpc.UpdateEmployee).Methods(http.MethodPut)                      				//http://localhost:8080/hostname/employee?Id=1&Name="Linus"&SecondName="Torvalds"&Surname="Benedict"&Position=developer&HireDate=2017-03-03&CompanyId=1
	r.HandleFunc(cnf.Path+"/employee/{employeeId}", rpc.GetEmployeeById).Methods(http.MethodGet)        				//http://localhost:8080/hostname/employee/1?Id=&Name="Linus"
	r.HandleFunc(cnf.Path+"/employee/{employeeId}", rpc.DeleteEmployee).Methods(http.MethodDelete)      				//http://localhost:8080/hostname/employee/1?Id=1
	r.HandleFunc(cnf.Path+"/employee/{employeeId}", rpc.UpdateEmployeeWithForm).Methods(http.MethodPut) 				//http://localhost:8080/hostname/employee/1?Id=1&Name="Linus"&SecondName="Torvalds"&Surname="Benedict"&Position=developer&HireDate=2017-03-03&CompanyId=1

	host := fmt.Sprintf("%s:%s", cnf.Host, cnf.Port)
	http.ListenAndServe(host, r)
}
