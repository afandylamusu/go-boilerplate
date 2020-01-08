package customer

import (
	"github.com/afandylamusu/moonlay.mcservice/models"
)

// Service represent the customer service
type Service interface {
	Fetch(query interface{}, args ...interface{}) ([]models.Customer, error)
}
