package customer

import (
	"github.com/afandylamusu/go-boilerplate/models"
)

// Service represent the customer service
type Service interface {
	Fetch(query interface{}, args ...interface{}) ([]models.Customer, error)
}
