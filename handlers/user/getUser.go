package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET user",
	})
}
