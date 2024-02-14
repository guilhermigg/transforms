package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "CREATE user",
	})
}
