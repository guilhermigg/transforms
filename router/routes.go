package router

import (
	handler "transforms/handlers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandlers()

	v1 := router.Group("/api/v1")
	{
		// User endpoints
		v1.GET("/user/:id", handler.GetUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.PATCH("/user/:id", handler.UpdateUserHandler)
		v1.DELETE("/user/:id", handler.DeleteUserHandler)
	}
}
