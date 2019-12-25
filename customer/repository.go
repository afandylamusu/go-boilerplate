package customer

import (
	"github.com/afandylamusu/stnkku.mdm/models"
)

// Repository represent the customer repository
type Repository interface {
	Fetch(query interface{}, args ...interface{}) ([]models.Customer, error)
	Store(m *models.Customer) error
}
