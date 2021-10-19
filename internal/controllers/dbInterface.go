package controllers

import "BookStore/internal/models"

// Repository is an interface that provides all required database-related methods.
type Repository interface {
	GetBooks() ([]models.BStore, error)
}
