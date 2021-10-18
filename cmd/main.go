package main

import (
	"BookStore/internal/app"
	"BookStore/internal/config"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// AppName is the name of application. We use it for logging messages
const AppName = "BookStore: "

func main() {
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "Start loading config...")
	c := config.AppConfig{}
	if err := c.Load(config.SERVICENAME); err != nil {
		logrus.Fatalf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Can't load config, missing environment variable. Error:", err.Error())
	}
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "Successfully load config")
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "Start initializing application...")
	router, err := app.Initialize(c)
	if err != nil {
		logrus.Fatalf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Can't initialize application. Error:", err.Error())
	}
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "Application has been initialized")
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "Usual router is listeting ...")
	err = http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		logrus.Fatalf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Service shutdown. Error:", err.Error())
	}
}
