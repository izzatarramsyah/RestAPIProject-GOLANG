package viewmodels

import "../models"

type Response struct {
	Code        string   	 	
	Message		string
	Object		[]models.Employee
}
