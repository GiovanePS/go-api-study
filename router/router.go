package router

import (
	gin "github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRouter(router)

	router.Run(":1024")
}
