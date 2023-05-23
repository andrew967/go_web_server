package migrations

import (
	"go_web_server/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func RunMigrations(db *gorm.DB) error {
	db.Logger.LogMode(logger.Info)

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Authorization{}); err != nil {
		return err
	}

	return nil

}
