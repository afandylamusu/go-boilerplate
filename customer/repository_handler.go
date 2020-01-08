package customer

import (
	"time"

	"github.com/afandylamusu/go-boilerplate/dbconn"
	"github.com/afandylamusu/go-boilerplate/models"
	uuid "github.com/satori/go.uuid"
)

// repositoryHandler the Customer Repository Handler
type repositoryHandler struct {
	conn *dbconn.DbConnection
}

// NewRepository to create new repo instance
func NewRepository(conn *dbconn.DbConnection) Repository {
	return &repositoryHandler{conn}
}

// Fetch customer from table
func (s *repositoryHandler) Fetch(offset int, limit int, query interface{}, args ...interface{}) ([]models.Customer, error) {
	var rows []models.Customer

	s.conn.GetDb().Where(query, args).Offset(offset).Limit(limit).Find(&rows)

	return rows, nil
}

// Store a customer
func (s *repositoryHandler) Store(id uuid.UUID, firstName string, lastName string, user string, t time.Time, tested bool) (*models.Customer, error) {

	m := &models.Customer{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		BaseModel: models.BaseModel{
			CreatedBy: user,
			CreatedAt: t,
			UpdatedAt: t,
			UpdatedBy: user,
			Deleted:   false,
			Tested:    tested,
		},
	}

	err := s.conn.GetDb().Create(m).Error
	if err != nil {
		return nil, err
	}

	mtrail := m.ToTrail()
	err = s.conn.GetDbTrail().Create(mtrail).Error

	if err != nil {
		panic(err)
	}

	return m, nil
}

// Update a modified customer
func (s *repositoryHandler) Update(m *models.Customer) {
	m.UpdatedBy = "user"
	m.UpdatedAt = time.Now()

	s.conn.GetDb().Save(&m)

	mtrail := m.ToTrail()
	s.conn.GetDbTrail().Create(&mtrail)
}
