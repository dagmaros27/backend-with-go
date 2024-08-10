package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"task_managment_api/data"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization format must be Bearer {token}"})
            return
        }

        tokenString := parts[1]

        
        token, err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{}, error) {

            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(data.JwtSecret), nil
        })
        
        
        if  err != nil || !token.Valid {
            
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }
        
        claims,ok := token.Claims.(jwt.MapClaims)

        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        }

        c.Set("userId", claims["userId"])
        c.Set("username", claims["username"])
        c.Set("role", claims["role"])
        c.Next()
    }
}



func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
        fmt.Println(role)
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}

