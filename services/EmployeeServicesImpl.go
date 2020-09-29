package services

import (
	"../models"
	"../interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/Sirupsen/logrus"
)

type EmployeeServicesImpl struct{
	Impl 		  interfaces.EmployeeRepository
	Log 		  *logrus.Logger
}

func (services *EmployeeServicesImpl) GetEmployeeServices(filter bson.M) *models.Response{
	services.Log.Info("Endpoint Hit: Get Service ")
	return services.Impl.GetEmployeeRepo(filter)
}

func (services *EmployeeServicesImpl) CreateEmployeeServices(emp models.Employee) *models.Response{
	services.Log.Info("Endpoint Hit: Create Service ")
	return services.Impl.CreateEmployeeRepo(emp)
}