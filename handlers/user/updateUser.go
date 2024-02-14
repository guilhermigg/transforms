package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "UPDATE user",
	})
}
