package handlers

import (
	"BookStore/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// AppName is the name of application. We use it for logging messages
const (
	AppName           = "BookStore: "
	ErrRequestBody    = "bad request body"
	ErrServerInternal = "internal server error"
	ErrPostForm       = "invalid post form"
)

// GetAllBooks returns all books in the DB
func (h *HTTPHandler) GetAllBooks(rw http.ResponseWriter, req *http.Request) {
	books, err := h.controller.GetAllBooks()
	if err != nil {
		http.Error(rw, ErrServerInternal, http.StatusInternalServerError)
		logrus.Errorf(AppName + "[" + time.Now().Format(time.RFC822) + "] " + fmt.Sprintf("can't send welcome message to the user with phone %s. Error:%s", req.FormValue("Phone"), err))
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(books); err != nil {
		logrus.Errorf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot encode messages! ", err.Error())
		http.Error(rw, ErrRequestBody, http.StatusBadRequest)
		return
	}
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "All books has been sent")
}

// AddBook tries to decode the json file. If success - add book to the DB.
func (h *HTTPHandler) AddBook(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		logrus.Warn(context.Background(), AppName+"["+time.Now().Format(time.RFC822)+"] "+ErrPostForm, nil)
		http.Error(rw, ErrPostForm, http.StatusBadRequest)
		return
	}
	var book models.BStore
	err = json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		logrus.Info(ErrRequestBody)
		http.Error(rw, ErrRequestBody, http.StatusBadRequest)
		return
	}
	err = checkBookDescription(book)
	if err != nil {
		logrus.Info(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.controller.AddBook(book)
	if err != nil {
		logrus.Error("cannot add book to the DB:", err)
		http.Error(rw, ErrServerInternal+" cannot add book", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "book has been saved")
}

// DeleteBookByID tries to decode the json file. If success - delete book from the DB.
func (h *HTTPHandler) DeleteBookByID(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		logrus.Warn(context.Background(), AppName+"["+time.Now().Format(time.RFC822)+"] "+ErrPostForm, nil)
		http.Error(rw, ErrPostForm, http.StatusBadRequest)
		return
	}
	var bookID int
	err = json.NewDecoder(req.Body).Decode(&bookID)
	if err != nil {
		logrus.Info(ErrRequestBody)
		http.Error(rw, ErrRequestBody, http.StatusBadRequest)
		return
	}
	if bookID < 0 {
		logrus.Warn(context.Background(), AppName+"["+time.Now().Format(time.RFC822)+"] "+"Wrong ID number", nil)
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = h.controller.DeleteBookByID(bookID)
	if err != nil {
		logrus.Error("cannot delete book from the DB by:", err)
		http.Error(rw, ErrServerInternal+" cannot delete book", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "book has been deleted")
}

// UpdateBookByID tries to decode the json file. If success - delete book from the DB.
func (h *HTTPHandler) UpdateBookByID(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		logrus.Warn(context.Background(), AppName+"["+time.Now().Format(time.RFC822)+"] "+ErrPostForm, nil)
		http.Error(rw, ErrPostForm, http.StatusBadRequest)
		return
	}
	var book models.BStore
	err = json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		logrus.Info(ErrRequestBody)
		http.Error(rw, ErrRequestBody, http.StatusBadRequest)
		return
	}
	err = checkBookDescription(book)
	if err != nil {
		logrus.Info(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.controller.UpdateBookByID(book)
	if err != nil {
		logrus.Error("cannot update book from the DB by:", err)
		http.Error(rw, ErrServerInternal+" cannot update book", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "book has been updated")
}

// FindBookBookByParameters tries to search book by special parameters: title, author, publication date
func (h *HTTPHandler) FindBookBookByParameters(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		logrus.Warn(context.Background(), AppName+"["+time.Now().Format(time.RFC822)+"] "+ErrPostForm, nil)
		http.Error(rw, ErrPostForm, http.StatusBadRequest)
		return
	}
	var title, author, publicationDate string
	if req.FormValue("title") != "" {
		title = req.FormValue("title")
	}
	if req.FormValue("author") != "" {
		author = req.FormValue("author")
	}
	if req.FormValue("publication_date") != "" {
		publicationDate = req.FormValue("publication_date")
	}
	books, err := h.controller.FindBookByParameters()
	if err != nil {
		logrus.Error("cannot update book from the DB by:", err)
		http.Error(rw, ErrServerInternal+" cannot update book", http.StatusInternalServerError)
		return
	}
	logrus.Info(title, author, publicationDate)
	rw.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(rw).Encode(books); err != nil {
		logrus.Errorf(AppName+"["+time.Now().Format(time.RFC822)+"] "+"Cannot encode messages! ", err.Error())
		http.Error(rw, ErrRequestBody, http.StatusBadRequest)
		return
	}
	logrus.Info(AppName + "[" + time.Now().Format(time.RFC822) + "] " + "All books has been sent")
}

func checkBookDescription(book models.BStore) error {
	switch {
	case book.ID < 0:
		return fmt.Errorf("id cannot be less than 0")
	case book.Author == "":
		return fmt.Errorf("'author' field cannot be empty")
	case book.Title == "":
		return fmt.Errorf("'title' field cannot be empty")
	default:
		return nil
	}
}
