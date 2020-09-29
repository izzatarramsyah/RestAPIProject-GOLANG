package services

import (
	"testing"
	"../models"
	"../mocks"
	"github.com/stretchr/testify/assert"
	"github.com/Sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateEmployeeRepo(t *testing.T){

	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}

	employee := models.Employee{ Id : 3, Name : "Thor", Nik : "113", Division : "IT Developer"}

	employeeRepository := new(mocks.EmployeeRepository)
	service := EmployeeServicesImpl{Impl : employeeRepository, Log : logger}
	actualResult := service.CreateEmployeeServices(employee)

	expectedResult := models.Wrap("200", "Success", nil)

	assert.Equal(t, expectedResult, actualResult)
}

func TestGetEmployeeRepo(t *testing.T){

	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}

	employeeRepository := new(mocks.EmployeeRepository)

	var employees = []models.Employee { 
		models.Employee {
			Id: 1, 
			Name: "Iron Man", 
			Nik: "111", 
			Division: "IT Developer",
		},
	}

	employeeRepository.On("GetEmployeeRepo", bson.M{"id": 1})
	service := EmployeeServicesImpl{Impl : employeeRepository, Log : logger}
	actualResult := service.GetEmployeeServices(bson.M{"id": 1})
	expectedResult := models.Wrap("200", "Success", employees)

	assert.Equal(t, expectedResult, actualResult)
}

