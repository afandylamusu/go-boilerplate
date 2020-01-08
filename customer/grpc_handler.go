package customer

import (
	context "context"

	"github.com/afandylamusu/go-boilerplate/dbconn"
)

// GrpcHandler the GRPC Handler
type GrpcHandler struct {
	Port string
	Db   *dbconn.DbConnection
}

// Find customer from grpc
func (s *GrpcHandler) Find(ctx context.Context, request *FindCustomerRequest) (*FindCustomerResponse, error) {
	// customerId := request.GetCustomerID()

	return &FindCustomerResponse{Success: true}, nil
}
