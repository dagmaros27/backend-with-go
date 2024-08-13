package infrastructure

import (
	"net/http"
	"task_managment_api/domain"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashingCost = 10
)

func HashPassword(password string) (string, domain.CustomError) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost )
	if err != nil {
		return "", domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: "Error while hashing password"}
	}
	return string(hashedPassword), domain.CustomError{}
}


func VerifyPassword(user domain.User, password string) domain.CustomError{
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return  domain.CustomError{ErrCode: http.StatusUnauthorized, ErrMessage: "Invalid username or password"}
	}
	return domain.CustomError{}
}