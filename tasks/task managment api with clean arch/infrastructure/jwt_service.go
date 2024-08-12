package infrastructure

import (
	"task_managment_api/domain"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func GenerateUserToken(user domain.User) (string, domain.CustomError) {{

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &domain.Claims{
		UserId:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(domain.JwtSecret))
	if err != nil {
		return "", domain.CustomError{ErrCode: http.StatusInternalServerError, ErrMessage: err.Error()}
	}

	return tokenString, domain.CustomError{}
}}


