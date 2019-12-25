package models

// Kurir is model
type Kurir struct {
	AuditModel
	ID        string `json:"id" gorm:"size:32;primary_key"`
	FirstName string `json:"first_name" gorm:"size:64;not null"`
	LastName  string `json:"last_name" gorm:"size:64"`
}

// ToTrail is to convert to KurirTrail
func (m *Kurir) ToTrail() *KurirTrail {
	return &KurirTrail{
		AuditModel: m.AuditModel,
		KurirID:    m.ID,
		FirstName:  m.FirstName,
		LastName:   m.LastName,
	}
}

// KurirTrail is model for Kurir history
type KurirTrail struct {
	AuditModel
	KurirID   string `json:"id" gorm:"size:32"`
	FirstName string `json:"first_name" gorm:"size:64;not null"`
	LastName  string `json:"last_name" gorm:"size:64"`
}
