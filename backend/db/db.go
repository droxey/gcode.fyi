package db

import (
	"github.com/droxey/gcode.fyi/backend/models"
	"github.com/droxey/gcode.fyi/backend/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./example.db")
	utils.CheckError(err)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Command{},
	)
}
