package customer

import (
	"time"

	"github.com/afandylamusu/go-boilerplate/models"
	uuid "github.com/satori/go.uuid"
)

// Repository represent the customer repository
type Repository interface {
	Fetch(offset int, limit int, query interface{}, args ...interface{}) ([]models.Customer, error)
	Store(id uuid.UUID, firstName string, lastName string, user string, t time.Time, tested bool) (*models.Customer, error)
	Update(m *models.Customer)
}
