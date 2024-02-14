package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func compareHashAndPassword(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fmt.Errorf("comparing passwords gone wrong: %+v", err.Error())
	}

	return nil
}

type LoginData struct { // from request
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct { // from database
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUserHandler(ctx *gin.Context) {
	var user User

	request := LoginUserRequest{}

	ctx.BindJSON(&request)

	loginData := LoginData{
		Email:    request.Email,
		Password: request.Password,
	}

	fmt.Println(loginData.Email)

	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		logger.Error("user not found: %+v", err.Error())
		sendError(ctx, "login invalid", http.StatusBadRequest)
		return
	}

	if err := compareHashAndPassword(user.Password, loginData.Password); err != nil {
		logger.Error("wrong password %v", err.Error())
		sendError(ctx, "login invalid", http.StatusBadRequest)
		return
	}

	sendSuccess(ctx, gin.H{
		"username": user.Username,
		"email":    user.Email,
	})
}
