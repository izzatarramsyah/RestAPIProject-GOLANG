package models

type Employee struct {
	Id          int 	`json:"id"`
	Name      	string  `json:"name"`
	Nik 		string 	`json:"nik"`
	Division	string  `json:"division"`
}

