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

// AddBook saves through the controller book to the DB
func (c *AppController) AddBook(book models.BStore) error {
	err := c.db.AddBook(book)
	return err
}

// DeleteBookByID deletes the book by id through the controller from the DB
func (c *AppController) DeleteBookByID(id int) error {
	err := c.db.DeleteBookByID(id)
	return err
}

// UpdateBookByID updates the book by id through the controller in the DB
func (c *AppController) UpdateBookByID(book models.BStore) error {
	err := c.db.UpdateBookByID(book)
	return err
}

// FindBookByParameters returns through the controller all books that had been find by special parameters from the DB
func (c *AppController) FindBookByParameters(parameters []string) ([]models.BStore, error) {
	books, err := c.db.FindBookByParameters(parameters)
	return books, err
}
