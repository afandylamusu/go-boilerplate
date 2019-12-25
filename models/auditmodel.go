package models

import "time"

// AuditModel is a model
type AuditModel struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by" gorm:"size:64"`
	UpdatedAt time.Time `json:"updated_at" gorm:"index"`
	UpdatedBy string    `json:"updated_by" gorm:"size:64"`
	Deleted   bool      `json:"deleted"`
}
