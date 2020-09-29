package main

import (
    "context"
	"os"
	"io"
	"net/http"
	"./controller"
	"./repository"
	"./services"
	"github.com/go-chi/chi"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
)

func main(){
	file,_ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger := &logrus.Logger{
        Out:   os.Stderr,
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%",
        },
	}
	logger.SetOutput(io.MultiWriter(file, os.Stdout))
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	if(err != nil){
		logger.Info("Connect Failed : ",err)
	}else{
		logger.Info("Connect Success")
	}

	err = client.Connect(context.Background())
	if err != nil {
        logger.Info("Connect Failed ", err)
	}
	
	repository := &repository.EmployeeRepositoryImpl{Conn : client, Log : logger}
	services   := &services.EmployeeServicesImpl{Impl : repository, Log : logger}
	controller := controller.EmployeeController{Impl : services, Log : logger}

	r := chi.NewRouter()
	r.Get("/GetAllEmployee", controller.GetAllEmployee)
	r.Post("/GetEmployee", controller.GetEmployee)
	r.Post("/CreateEmployee", controller.CreateEmployee)

	http.ListenAndServe(":8083", r)

	
}



