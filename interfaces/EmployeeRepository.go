package interfaces

import(
	"../models"
	"go.mongodb.org/mongo-driver/bson"
)

type EmployeeRepository interface{

	GetEmployeeRepo(filter bson.M) *models.Response
	CreateEmployeeRepo(emp models.Employee) *models.Response

}