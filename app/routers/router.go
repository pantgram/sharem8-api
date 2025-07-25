package routers

import (
	"sharem8-api/app/controllers/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users",users.RetrieveUsers)
	router.POST("/users",users.RetrieveUsers)
	return router
}