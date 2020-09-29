package interfaces

import(
	"../models"
	"go.mongodb.org/mongo-driver/bson"

)

type EmployeeServices interface{

	GetEmployeeServices(filter bson.M) *models.Response
	CreateEmployeeServices(emp models.Employee) *models.Response

}