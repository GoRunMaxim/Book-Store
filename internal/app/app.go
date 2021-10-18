package app

import (
	"BookStore/internal/config"
	"BookStore/internal/controllers"
	"BookStore/internal/handlers"
	"BookStore/internal/postgreSql"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	// ApplicationName is the name of app. Use it for logrus messages
	ApplicationName = "BookStore: "
)

// Initialize performs required application configuration.
func Initialize(cfg config.AppConfig) (externalRouter *mux.Router, err error) {
	err = initializeLogger(cfg.LogConfig)
	if err != nil {
		logrus.Errorf(ApplicationName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot initialize Logger: ", err.Error())
		return nil, err
	}

	controller, err := initializeAPIController(cfg)
	if err != nil {
		logrus.Errorf(ApplicationName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot initialize APIController: ", err.Error())
		return nil, err
	}

	router := initializeRouter(controller)

	return router, nil
}

// initializeAPIController performs required controller configuration.
func initializeAPIController(cfg config.AppConfig) (handlers.Controller, error) {
	DB, err := postgreSql.New(cfg.DbConfig)
	if err != nil {
		logrus.Errorf(ApplicationName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot create DB: ", err.Error())
		return nil, err
	}
	controller := controllers.New(DB)
	return controller, nil
}

func initializeRouter(controller handlers.Controller) *mux.Router {
	var router = mux.NewRouter()
	// var handler = handlers.NewHTTPHandler(controller)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {}).Methods("GET")
	return router
}

func initializeLogger(cfg config.LogConfig) error {
	if cfg.WriteToFile {
		f, err := os.OpenFile(cfg.Filepath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	return nil
}
