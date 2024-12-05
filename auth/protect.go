package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Protect(signature []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 	aud := claims["aud"].(string)
		// 	c.Set("aud", aud)
		// }
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			aud := claims["aud"]
			// If "aud" is not a string, it could be an array, so handle that case
			switch v := aud.(type) {
			case string:
				c.Set("aud", v) // If "aud" is a string, store it
			case []interface{}:
				if len(v) > 0 {
					c.Set("aud", fmt.Sprintf("%v", v[0])) // Store the first element if it's an array
				}
			}
		}

		c.Next()
	}
}
