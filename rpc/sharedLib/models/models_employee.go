package models

type Employee struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name"`
	SecondName string `json:"secondName,omitempty"`
	Surname    string `json:"surname,omitempty"`
	HireDate   string `json:"hireDate,omitempty"`
	Position   string `json:"position,omitempty"`
	CompanyId  int64  `json:"companyId,omitempty"`
}
