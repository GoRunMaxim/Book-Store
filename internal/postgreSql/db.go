package postgreSql

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

func (D *DB) GetAllBooks() ([]models.BStore, error) {
	panic("implement me")
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
