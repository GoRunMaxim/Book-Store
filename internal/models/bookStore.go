package models

import "time"

// BStore represents information about Book Storage struct
type BStore struct {
	ID int
	Title string
	Author string
	PublicDate time.Time
	PagesAmount int
	CreatedTime time.Time
	UpdatedTime time.Time
}
