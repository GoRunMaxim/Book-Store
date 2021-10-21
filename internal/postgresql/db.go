package postgresql

import (
	"BookStore/internal/config"
	"BookStore/internal/models"
	"database/sql"
	"time"
)

// DB is the repository, with all the methods that are required to get info from the db
type DB struct {
	db *sql.DB
}

// FindBookByParameters searches in the DB books by special parameters
func (d *DB) FindBookByParameters([]string) ([]models.BStore, error) {
	return nil, nil
}

// UpdateBookByID updates book by ID in the DB
func (d *DB) UpdateBookByID(models.BStore) error {
	return nil
}

// DeleteBookByID deletes book by ID from DB
func (d *DB) DeleteBookByID(int) error {
	return nil
}

// GetBooks return all book in the DB
func (d *DB) GetBooks() ([]models.BStore, error) {
	return nil, nil
}

// AddBook save book in the DB
func (d *DB) AddBook(models.BStore) error {
	return nil
}

// New returns new DB repository
func New(cfg config.DatabaseConfig) (*DB, error) {
	sqlDB, err := sql.Open(cfg.Dialect, cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(5 * time.Second)

	return &DB{db: sqlDB}, nil
}
