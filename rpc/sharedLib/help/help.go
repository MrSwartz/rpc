package help

import (
	"fmt"
	"net/http"
	"os"
	"rpc/sharedLib/models"
	"strconv"
)

func GetId(r *http.Request) int64 {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in GetId : %v\n", err)
	}

	id, err := strconv.ParseInt(r.Form["Id"][0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in GetId : %v \n", err)
	}

	return id
}

func ParseURLArgsCompany(r *http.Request) (cmp models.Company) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in ParseURLArgsCompany : %v \n", err)
	}

	id, err := strconv.ParseInt(r.Form["Id"][0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in ParseURLArgsCompany : %v \n", err)
	}

	cmp.Id = id
	cmp.Name = r.Form["Name"][0]
	cmp.LegalForm = r.Form["LegalForm"][0]

	return cmp
}

func ParseURLArgsEmployee(r *http.Request) (emp models.Employee) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in ParseURLArgsEmployee : %v\n", err)
	}

	id, err := strconv.ParseInt(r.Form["Id"][0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in ParseURLArgsEmployee : %v\n", err)
	}

	CompanyId, err := strconv.ParseInt(r.Form["CompanyId"][0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in ParseURLArgsEmployee : %v\n", err)
	}

	emp.Id = id
	emp.CompanyId = CompanyId
	emp.Name = r.Form["Name"][0]
	emp.HireDate = r.Form["HireDate"][0]
	emp.Position = r.Form["Position"][0]
	emp.SecondName = r.Form["SecondName"][0]
	emp.Surname = r.Form["Surname"][0]

	return emp
}
