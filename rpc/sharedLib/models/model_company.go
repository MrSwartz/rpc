package models

type Company struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	LegalForm string `json:"legalForm,omitempty"`
}
