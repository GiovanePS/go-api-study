package handler

import (
	"github.com/GiovanePS/go-api-study/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()
}
