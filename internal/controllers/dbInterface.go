package controllers

import "BookStore/internal/models"

// Repository is an interface that provides all required database-related methods.
type Repository interface {
	AddBook(models.BStore) error
	GetBooks() ([]models.BStore, error)
	DeleteBookByID(int) error
	UpdateBookByID(models.BStore) error
}
