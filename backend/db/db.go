package db

import (
	"github.com/droxey/gcode.fyi/backend/models"
	"github.com/droxey/gcode.fyi/backend/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New TODO...
func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gcode.db")
	utils.CheckError(err)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	db.DropTable(&models.Command{})
	db.DropTable(&models.Firmware{})
	db.DropTable(&models.Feature{})
	db.DropTable(&models.Parameter{})
	return db
}

// AutoMigrate TODO...
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Command{},
		&models.Firmware{},
		&models.Feature{},
		&models.Parameter{},
	)
}
