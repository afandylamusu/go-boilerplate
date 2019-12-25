package customer

import (
	"github.com/afandylamusu/stnkku.mdm/dbconn"
	"github.com/afandylamusu/stnkku.mdm/models"
)

// repositoryHandler the Customer Repository Handler
type repositoryHandler struct {
	conn *dbconn.Connection
}

// NewRepository to create new repo instance
func NewRepository(conn *dbconn.Connection) *repositoryHandler {
	return &repositoryHandler{conn}
}

// Fetch customer from table
func (s *repositoryHandler) Fetch(query interface{}, args ...interface{}) ([]models.Customer, error) {
	var rows []models.Customer

	s.conn.Db.Where(query, args).Find(&rows)

	return rows, nil
}

// Store a customer
func (s *repositoryHandler) Store(m *models.Customer) error {

	s.conn.Db.Create(m)
	s.conn.DbTrail.Create(m.ToTrail())

	return nil
}
