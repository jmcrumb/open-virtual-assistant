package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/auth"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= len(BEARER_SCHEMA) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenString := authHeader[len(BEARER_SCHEMA):] // ayo, runtime error: slice bounds out of range [6:0]
		tokenString = strings.TrimSpace(tokenString)

		token, err := auth.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			ctx := context.WithValue(c.Request.Context(), "account_id", claims["id"])
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
