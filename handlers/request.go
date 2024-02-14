package handler

import (
	"fmt"
	"regexp"
)

func paramIsRequiredError(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Username == "" {
		return paramIsRequiredError("Username", "string")
	}

	if r.Email == "" {
		return paramIsRequiredError("Email", "string")
	} else if !isValidEmail(r.Email) {
		return fmt.Errorf("Invalid email format")
	}

	if r.Password == "" {
		return paramIsRequiredError("Password", "string")
	} else if len(r.Password) < 8 {
		return fmt.Errorf("Password must be at least 8 characters long")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Email == "" {
		return paramIsRequiredError("Email", "string")
	}

	if r.Password == "" {
		return paramIsRequiredError("Password", "string")
	}

	return nil
}
