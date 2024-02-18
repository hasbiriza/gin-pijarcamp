package helpers

import (
	"latihan_git/src/config"
	"latihan_git/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Month{})
}
