package controllers

import (
	"BookStore/internal/models"
)

// AppName is the name of application. We use it for logging messages
const AppName = "BookStore: "

// GetAllBooks returns through the controller all books from the DB
func (c *AppController) GetAllBooks() ([]models.BStore, error) {
	books, err := c.db.GetBooks()
	return books, err
}

// AddBook save through the controller book to the DB
func (c *AppController) AddBook(book models.BStore) error {
	err := c.db.AddBook(book)
	return err
}

// DeleteBookByID delete the book by id through the controller from the DB
func (c *AppController) DeleteBookByID(id int) error {
	err := c.db.DeleteBookByID(id)
	return err
}
