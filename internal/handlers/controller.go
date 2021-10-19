package handlers

import "BookStore/internal/models"

// Controller is an interface with all application related controller methods.
type Controller interface {
	GetAllBooks() ([]models.BStore, error)
}
