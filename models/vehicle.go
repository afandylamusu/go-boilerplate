package models

import "time"

// Vehicle is a model
type Vehicle struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	Deleted   bool      `json:"deleted"`
}
