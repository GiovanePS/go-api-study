package router

import (
	gin "github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	initializeRouter(router)

	router.Run(":1024")
}
