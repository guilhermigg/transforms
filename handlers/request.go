package handler

import (
	"fmt"
	"regexp"
)
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

