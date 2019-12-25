package dbconn

import (
	"fmt"
	"log"

	"github.com/afandylamusu/stnkku.mdm/models"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Connection for gorm.DB
type Connection struct {
	Db      *gorm.DB
	DbTrail *gorm.DB
}

// Open database connection
func (c *Connection) Open() {
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

	c.Db = db

	dbtrail, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v_trail password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass))
	if err != nil {
		log.Fatalf("failed to established db connection: %v", err)
	}

	c.DbTrail = dbtrail

	c.Migrate()
}

// Close the active connection
func (c *Connection) Close() {
	c.Db.Close()
	c.DbTrail.Close()
}

// Migrate database
func (c *Connection) Migrate() {
	c.Db.AutoMigrate(&models.Customer{}, &models.Kurir{}, &models.Vehicle{})

	c.DbTrail.AutoMigrate(&models.CustomerTrail{}, &models.KurirTrail{}, &models.VehicleTrail{})
}
