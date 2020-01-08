package dbconn

import (
	"fmt"
	"log"

	"github.com/afandylamusu/moonlay.mcservice/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DbConnection for gorm.DB
type DbConnection struct {
	db      *gorm.DB
	dbTrail *gorm.DB
}

// NewDbConnection create DbConnection instance from existing db gorm
func NewDbConnection(db *gorm.DB, dbtrail *gorm.DB) *DbConnection {
	return &DbConnection{db, dbtrail}
}

// GetDb db reference of DB
func (c *DbConnection) GetDb() *gorm.DB {
	return c.db
}

// GetDbTrail db reference of DB
func (c *DbConnection) GetDbTrail() *gorm.DB {
	return c.dbTrail
}

// Open database connection
func (c *DbConnection) Open() {
	isLocal := viper.GetString("env") == "local"
	var dbHost, dbPort, dbUser, dbPass, dbName string

	if isLocal {
		dbHost = viper.GetString(`database-local.host`)
		dbPort = viper.GetString(`database-local.port`)
		dbUser = viper.GetString(`database-local.user`)
		dbPass = viper.GetString(`database-local.pass`)
		dbName = viper.GetString(`database-local.name`)
	} else {
		dbHost = viper.GetString(`database.host`)
		dbPort = viper.GetString(`database.port`)
		dbUser = viper.GetString(`database.user`)
		dbPass = viper.GetString(`database.pass`)
		dbName = viper.GetString(`database.name`)
	}

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		log.Fatalf("failed to established db connection: %v", err)
	}

	c.db = db

	dbtrail, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v_trail password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		log.Fatalf("failed to established db connection: %v", err)
	}

	c.dbTrail = dbtrail

	c.Migrate()
}

// Close the active connection
func (c *DbConnection) Close() {
	err := c.GetDb().Close()
	if err != nil {
		panic(err)
	}

	err = c.GetDbTrail().Close()
	if err != nil {
		panic(err)
	}
}

// Migrate database
func (c *DbConnection) Migrate() {
	c.GetDb().AutoMigrate(&models.Customer{})

	c.GetDbTrail().AutoMigrate(&models.CustomerTrail{})
}
