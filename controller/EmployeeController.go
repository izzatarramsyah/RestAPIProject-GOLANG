package controller

import (
	"../models"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/Sirupsen/logrus"
)

type EmployeeController struct{
	Impl		  interfaces.EmployeeServices
	Log 		  *logrus.Logger
}

func (controller *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request){
	controller.Log.Info("Endpoint Hit: Get Controller")
	
	var employee models.Employee
	var response *models.Response

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		controller.Log.Error(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response = controller.Impl.GetEmployeeServices(bson.M{"id": employee.Id})
	json.NewEncoder(w).Encode(response)

}

func (controller *EmployeeController) GetAllEmployee(w http.ResponseWriter, r *http.Request){
	controller.Log.Info("Endpoint Hit: Get Controller")
	
	var employee models.Employee
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		controller.Log.Error(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(controller.Impl.GetEmployeeServices(bson.M{}))

}

func (controller *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request){
	controller.Log.Info("Endpoint Hit: Create Controller")

	var employee models.Employee
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		controller.Log.Error(w, "Kindly enter data with the event Consta and Value only in order to update")
	}

	jsonString := string(reqBody)
	json.Unmarshal([]byte(jsonString), &employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(controller.Impl.CreateEmployeeServices(employee))
}