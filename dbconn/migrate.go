package dbconn

import (
	"github.com/afandylamusu/stnkku.mdm/models"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Customer{}, &models.Kurir{}, &models.Vehicle{})
}
