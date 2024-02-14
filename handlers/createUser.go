package handler

import (
	"net/http"
	"transforms/schemas"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashed)
}

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	encryptedPassword := hashPassword(request.Password)

	user := schemas.User{
		Username: request.Username,
		Email:    request.Email,
		Password: encryptedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Error("creating user failed: %v", err)
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})
}
