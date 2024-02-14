package handler

import (
	"transforms/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandlers() {
	logger = config.GetLogger("handler")
	db = config.GetDatabase()
}
