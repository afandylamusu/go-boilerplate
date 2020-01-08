package models

import uuid "github.com/satori/go.uuid"

// Customer is a customer
type Customer struct {
	ID        uuid.UUID `json:"id" gorm:"size:32;primary_key"`
	FirstName string    `json:"first_name" gorm:"size:64;not null"`
	LastName  string    `json:"last_name" gorm:"size:64"`
	BaseModel
}

// TableName of customer entity
func (c *Customer) TableName() string {
	return "customers"
}

// CustomerTrail is a customer history model
type CustomerTrail struct {
	CustomerID uuid.UUID `json:"id" gorm:"size:32"`
	FirstName  string    `json:"first_name" gorm:"size:64;not null"`
	LastName   string    `json:"last_name" gorm:"size:64"`
	BaseModel
}

// ToTrail is convert Customer to CustomerTrail
func (c *Customer) ToTrail() *CustomerTrail {
	return &CustomerTrail{
		CustomerID: c.ID,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		BaseModel:  c.BaseModel}
}

// BeforeSave customer record
// func (c *Customer) BeforeSave() (err error) {
// 	c.UpdatedAt = time.Now()
// 	if !c.IsValid() {
// 		err = errors.New("can't save invalid data")
// 	}
// 	return
// }

// BeforeCreate customer record
// func (c *Customer) BeforeCreate() (err error) {
// 	c.CreatedAt = time.Now()
// 	c.UpdatedAt = c.CreatedAt
// 	if !c.IsValid() {
// 		err = errors.New("can't save invalid data")
// 	}
// 	return
// }
