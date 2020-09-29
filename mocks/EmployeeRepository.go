package mocks

import (
	"../models"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type EmployeeRepository struct{
	mock.Mock
}

func (m * EmployeeRepository) GetEmployeeRepo(filter bson.M) *models.Response{
	// ret := m.Called(filter)

	// var r0 []models.Employee
	// if rf, ok := ret.Get(0).(func(bson.M) []models.Employee); ok {
	// 	r0 = rf(filter)
	// } else {
	// 	r0 = ret.Get(0).([]models.Employee)
	// }

	var employees = []models.Employee { 
		models.Employee {
			Id: 1, 
			Name: "Iron Man", 
			Nik: "111", 
			Division: "IT Developer",
		},
	}

	commonResponse := models.Wrap("200", "Success", employees)
	return commonResponse
}

func (m * EmployeeRepository) CreateEmployeeRepo(emp models.Employee) *models.Response{
	commonResponse := models.Wrap("200", "Success", nil)
	return commonResponse
}