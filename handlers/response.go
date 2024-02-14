package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, msg string, code int) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}
