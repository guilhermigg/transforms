package config

import (
	"fmt"
	"os"
	"transforms/schemas"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database")

	godotenv.Load("development.env")

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error("database connection error: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Form{})
	err = db.AutoMigrate(&schemas.User{})

	if err != nil {
		logger.Error("automigrate error: %s", err)
		return nil, err
	}

	return db, nil
}
