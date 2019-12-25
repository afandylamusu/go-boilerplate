package customer

import (
	context "context"

	"github.com/afandylamusu/stnkku.mdm/models"
	"github.com/jinzhu/gorm"
)

// ServiceHandler the GRPC Handler
type ServiceHandler struct {
	Port string
	Db   *gorm.DB
}

// Find customer from grpc
func (s *ServiceHandler) Find(ctx context.Context, request *FindCustomerRequest) (*FindCustomerResponse, error) {
	// customerId := request.GetCustomerID()

	return &FindCustomerResponse{Success: true}, nil
}

// Fetch customer from table
func (s *ServiceHandler) Fetch(query interface{}, args ...interface{}) ([]models.Customer, error) {
	var rows []models.Customer

	s.Db.Where(query, args).Find(&rows)

	return rows, nil
}
