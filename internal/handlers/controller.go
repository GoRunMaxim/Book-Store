package handlers

import "BookStore/internal/models"

// Controller is an interface with all application related controller methods.
type Controller interface {
	GetAllBooks() ([]models.BStore, error)
	AddBook(models.BStore) error
	DeleteBookByID(int) error
	UpdateBookByID(models.BStore) error
	FindBookByParameters() ([]models.BStore, error)
}
