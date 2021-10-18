package postgreSql

import (
	"Itechart/BookStore/Book-Store/internal/config"
	"database/sql"
	"time"
)

// DB is the repository, with all the methods that are required to get info from the db
type DB struct {
	db *sql.DB
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
