package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, msg string, code int) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
