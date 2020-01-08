package customer

// https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b

import (
	"database/sql"
	"testing"
	"time"

	"github.com/afandylamusu/go-boilerplate/dbconn"
	"github.com/afandylamusu/go-boilerplate/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type Suite struct {
	suite.Suite
	DB      *gorm.DB
	DBTrail *gorm.DB
	mock    sqlmock.Sqlmock

	repository Repository
	customer   *models.Customer
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)

	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = NewRepository(dbconn.NewDbConnection(s.DB, s.DB))
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Repository_Fetch() {
}

func (s *Suite) Test_Repository_Store() {
	var (
		id        = uuid.NewV4()
		firstName = "test-first-name"
		lastName  = "test-last-name"
		now       = time.Now()
		user      = "afandy"
		tested    = false
		// deleted   = false
	)

	// s.mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO customers (id, first_name, last_name, created_at, created_by, updated_at, updated_by, deleted, tested)
	// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING customers.id`)).
	// 	WithArgs(id, firstName, lastName, now.Format(timeFormat), user, now.Format(timeFormat), user, deleted, tested).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"id"}).AddRow(id.String()))

	ncustomer, err := s.repository.Store(id, firstName, lastName, user, now, tested)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), ncustomer)
}

func (s *Suite) Test_Repository_Update() {

}
