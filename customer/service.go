package customer

import (
	"github.com/afandylamusu/go-boilerplate/models"
)

// Service represent the customer service
type Service interface {
	CreateCustomer(firstName string, lastName string) (*models.Customer, error)
}
