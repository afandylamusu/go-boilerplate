package models

// Vehicle is a model
type Vehicle struct {
	AuditModel
	ID   string `json:"id" gorm:"size:32;primary_key"`
	Name string `json:"name" gorm:"size:128"`
}

// ToTrail is to convert to VehicleTrail
func (m *Vehicle) ToTrail() *VehicleTrail {
	return &VehicleTrail{
		AuditModel: m.AuditModel,
		VehicleID:  m.ID,
		Name:       m.Name,
	}
}

// VehicleTrail is a model for Vehicle history
type VehicleTrail struct {
	AuditModel
	VehicleID string `json:"id" gorm:"size:32"`
	Name      string `json:"name" gorm:"size:128"`
}
