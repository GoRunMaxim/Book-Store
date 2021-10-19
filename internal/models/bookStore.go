package models

import "time"

// BStore represents information about Book Storage struct
type BStore struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublicDate  time.Time `json:"public_date"`
	PagesAmount int       `json:"pages_amount"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
