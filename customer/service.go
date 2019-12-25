package customer

import (
	"github.com/afandylamusu/stnkku.mdm/models"
)

// Service represent the customer service
type Service interface {
	Fetch(query interface{}, args ...interface{}) ([]models.Customer, string, error)
}
