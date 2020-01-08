package models

import "time"

// BaseModel is a model
type BaseModel struct {
	// Tested indicate the data is testing
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by" gorm:"size:64"`

	UpdatedAt time.Time `json:"updated_at" gorm:"index"`
	UpdatedBy string    `json:"updated_by" gorm:"size:64"`

	Deleted bool `json:"deleted"`
	Tested  bool
}

// MakeTested flagging the record as data testing
func (m *BaseModel) MakeTested() {
	m.Tested = true
}

// MakeDelete flagging the record to be delete
func (m *BaseModel) MakeDelete() {
	m.Deleted = true
}

// IsValid validate the record
func (m *BaseModel) IsValid() bool {
	return true
}
