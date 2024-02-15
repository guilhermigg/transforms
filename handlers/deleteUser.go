package handler

import (
	"net/http"
	"strconv"
	"transforms/schemas"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	var user schemas.User
	userIdParam := ctx.Param("id")

	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		sendError(ctx, "invalid user id", http.StatusBadRequest)
		return
	}

	user.ID = uint(userId)

	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, "could not delete this user", http.StatusBadRequest)
		return
	}

	sendSuccess(ctx, gin.H{
		"msg": "user successfully deleted",
	})
}
