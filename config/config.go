package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	logger := GetLogger("Config")
	var err error

	db, err = InitializeDatabase()

	if err != nil {
		logger.Error("initialize database")
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
