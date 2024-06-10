package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"74edf612f393b4eb01fbc2c29dd96671": "12345",
	"d88b4b1e77c70ba780b56032db1c259b": "98765",
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")[1]
		if userID, exists := users[token]; exists {
			c.Set("userID", userID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
