package customer

import (
	"time"

	"github.com/afandylamusu/go-boilerplate/models"
	uuid "github.com/satori/go.uuid"
)

type serviceHandler struct {
	customerRepo Repository
}

// NewService for creating Service instance
func NewService(repo Repository) Service {
	return &serviceHandler{repo}
}

// CreateCustomer
func (s *serviceHandler) CreateCustomer(firstName string, lastName string) (*models.Customer, error) {
	newCustomerID := uuid.NewV4()

	customer, err := s.customerRepo.Store(newCustomerID, firstName, lastName, "system", time.Now(), false)
	if err != nil {
		panic(err)
	}

	return customer, err
}
