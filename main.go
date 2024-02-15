package main

import (
	"github.com/GiovanePS/go-api-study/config"
	"github.com/GiovanePS/go-api-study/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Init()
}
