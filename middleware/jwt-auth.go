package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/example/gin_framework/helper"
	"github.com/example/gin_framework/service"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		BEARER := "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER):]

		token, err := service.NewJwtService().ValidateJWT(tokenString)
		fmt.Println("this is token string:", err)
		if token.Claims != nil {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("ID", claims["ID"])
			log.Println("Username", claims["Username"])
			log.Println("ExpiredAt", claims["ExpiredAt"])
			log.Println("IssuedAt", claims["IssuedAt"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrorResponse(err))
		}
	}
}
