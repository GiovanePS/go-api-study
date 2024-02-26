package config

import (
	"fmt"
	"os"

	"github.com/GiovanePS/go-api-study/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("initializing error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("migrations error: %v", err)
	}

	return db, nil
}
