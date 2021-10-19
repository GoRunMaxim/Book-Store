package controllers

import (
	"BookStore/internal/models"
)

// AppName is the name of application. We use it for logging messages
const AppName = "BookStore: "

func (c *AppController) GetAllBooks() ([]models.BStore, error) {
	books, err := c.db.GetBooks()
	return books, err
}
