package models

// Customer is a customer
type Customer struct {
	AuditModel
	ID        string `json:"id" gorm:"size:32;primary_key"`
	FirstName string `json:"first_name" gorm:"size:64;not null"`
	LastName  string `json:"last_name" gorm:"size:64"`
}

// ToTrail is convert Customer to CustomerTrail
func (c *Customer) ToTrail() *CustomerTrail {
	return &CustomerTrail{AuditModel: c.AuditModel,
		CustomerID: c.ID,
		FirstName:  c.FirstName,
		LastName:   c.LastName}
}

// CustomerTrail is a customer history model
type CustomerTrail struct {
	AuditModel
	CustomerID string `json:"id" gorm:"size:32"`
	FirstName  string `json:"first_name" gorm:"size:64;not null"`
	LastName   string `json:"last_name" gorm:"size:64"`
}
