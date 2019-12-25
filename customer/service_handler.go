package customer

import (
	context "context"

	"github.com/afandylamusu/stnkku.mdm/dbconn"
)

// ServiceHandler the GRPC Handler
type ServiceHandler struct {
	Port string
	Db   *dbconn.Connection
}

// Find customer from grpc
func (s *ServiceHandler) Find(ctx context.Context, request *FindCustomerRequest) (*FindCustomerResponse, error) {
	// customerId := request.GetCustomerID()

	return &FindCustomerResponse{Success: true}, nil
}
