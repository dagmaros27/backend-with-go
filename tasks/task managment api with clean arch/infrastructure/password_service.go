package infrastructure

import (
	"net/http"
	"task_managment_api/domain"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashingCost = 10
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost )
	if err != nil {
		return "",err
	}
	return string(hashedPassword), nil
}


func VerifyPassword(user domain.User, password string)domain.CustomError{
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return  domain.CustomError{ErrCode: http.StatusUnauthorized, ErrMessage: "Invalid username or password"}
	}
	return domain.CustomError{}
}