package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// AppName is the name of application. We use it for logging messages
const AppName = "BookStore: "

// GetAllBooks returns all books in the DB
func (h *HTTPHandler) GetAllBooks(rw http.ResponseWriter, req *http.Request) {
	books, err := h.controller.GetAllBooks()
	if err != nil {
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		logrus.Errorf(AppName + "[" + time.Now().Format(time.RFC822) + "] " + fmt.Sprintf("can't send welcome message to the user with phone %s. Error:%s", req.FormValue("Phone"), err))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(books); err != nil {
		logrus.Errorf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot encode messages! ", err.Error())
	}
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + fmt.Sprintf("All books has been sent"))
}
