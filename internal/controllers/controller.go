package controllers

// AppController is used to handle the business logic of the application
type AppController struct {
	db Repository
}

// New returns new application controller with provided database.
func New(db Repository) *AppController {
	return &AppController{
		db: db,
	}
}
