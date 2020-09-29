package repository

import (
	"context"
	"../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Sirupsen/logrus"
)										

var message string

type EmployeeRepositoryImpl struct{
	Conn 		  *mongo.Client
	Log 		  *logrus.Logger
}

func (repository *EmployeeRepositoryImpl) GetEmployeeRepo(filter bson.M) *models.Response{
	repository.Log.Info("Endpoint Hit: Get Repository ")

	col := repository.Conn.Database("XLAxiata").Collection("Employee")
	cur, err := col.Find(context.TODO(), filter)

	if err != nil {
		repository.Log.Error("Error get employee ", err)
		message,_ = "Failed. Error get employee ", err
		commonResponse := models.Wrap("500", message, nil)
		return commonResponse
	}

	var result []models.Employee

	for cur.Next(context.TODO()) {
        var row models.Employee
        err = cur.Decode(&row)
        if err != nil {
			repository.Log.Error("Error decode employee ", err)
        }
        result = append(result, row)
	}

	commonResponse := models.Wrap("200","Success" , result)

	return commonResponse
}

func (repository *EmployeeRepositoryImpl) CreateEmployeeRepo(emp models.Employee) *models.Response{
	repository.Log.Info("Endpoint Hit: Create Repository ")

	commonResponse := models.Wrap("200", "Success", nil)

	col := repository.Conn.Database("XLAxiata").Collection("Employee")
	_,err := col.InsertOne(context.TODO(), emp)
	
	if err != nil{
		repository.Log.Error("Error insert employee ", err)
		message,_ = "Failed. Error create employee ", err
		commonResponse = models.Wrap("500", message, nil)
	}
	
	return commonResponse
}

