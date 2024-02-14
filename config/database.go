package config

import (
	"os"
	"transforms/schemas"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	logger := GetLogger("Database")

	godotenv.Load(".env")

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port + " sslmode=disable TimeZone=America/Sao_Paulo"
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
