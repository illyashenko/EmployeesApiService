package models

type Employee struct {
	DomainId string `json:"domainId"`
	Fio      string `json:"fio"`
	PassId   string `json:"passId"`
	Active   bool   `json:"active"`
	DateIn   string `json:"dateIn"`
	DateOut  string `json:"dateOut"`
}
