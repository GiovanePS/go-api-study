package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("dotenv error: %v", err.Error())
	}

	db, err = InitializePostgreSQL()
	if err != nil {
		return fmt.Errorf("postgres error: %v", err.Error())
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
