package controller

import (
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi"
	"../models"
	"../mocks"
	"../viewmodels"
	"github.com/stretchr/testify/assert"
	"github.com/Sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateEmployee(t *testing.T){

	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}

	employeeServices := new(mocks.EmployeeServices)
	employee := models.Employee{ Id : 1, Name : "Iron Man", Nik : "111", Division : "IT Developer"}
	employeeServices.On("CreateEmployeeServices", employee)

	controller := EmployeeController{Impl : employeeServices, Log : logger}

	// call the code we are testing
	req := httptest.NewRequest("POST", "http://localhost:8083/CreateEmployee", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/CreateEmployee", controller.CreateEmployee)
	r.ServeHTTP(w, req)

	expectedResult := 
		viewmodels.Response {
			Code: "200", 
			Message: "Success", 
			Object: nil, 
		}

	actualResult := viewmodels.Response{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}

func TestGetEmployee(t *testing.T){

	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}

	employees := []models.Employee { 
		models.Employee {
			Id: 1, 
			Name: "Iron Man", 
			Nik: "111", 
			Division: "IT Developer",
		},
	}

	employeeServices := new(mocks.EmployeeServices)
	employeeServices.On("GetEmployeeServices", bson.M{"id":1})
	controller := EmployeeController{Impl : employeeServices, Log : logger}

	// call the code we are testing
	req := httptest.NewRequest("POST", "http://localhost:8083/GetEmployee", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/GetEmployee", controller.GetEmployee)
	r.ServeHTTP(w, req)

	expectedResult := 
	viewmodels.Response {
		Code: "200", 
		Message: "Success", 
		Object: employees, 
	}

	actualResult := viewmodels.Response{}
	json.NewDecoder(w.Body).Decode(&actualResult)

		// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}

func TestGetAllEmployee(t *testing.T){

	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}

	employees := []models.Employee { 
		models.Employee {
			Id: 1, 
			Name: "Iron Man", 
			Nik: "111", 
			Division: "IT Developer",
		},
	}

	employeeServices := new(mocks.EmployeeServices)
	employeeServices.On("GetEmployeeServices", bson.M{})
	controller := EmployeeController{Impl : employeeServices, Log : logger}

	// call the code we are testing
	req := httptest.NewRequest("POST", "http://localhost:8083/GetAllEmployee", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/GetAllEmployee", controller.GetAllEmployee)
	r.ServeHTTP(w, req)

	expectedResult := 
	viewmodels.Response {
		Code: "200", 
		Message: "Success", 
		Object: employees, 
	}

	actualResult := viewmodels.Response{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}