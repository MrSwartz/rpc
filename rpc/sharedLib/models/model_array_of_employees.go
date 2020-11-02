package models

type ArrayOfEmployees struct {
	Employees []Employee
}

type MyNotFavouriteUniqueStruct struct {
	Code int32  `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
	Payload
}

type Payload struct {
	Id  int64
	Emp []Employee
}
