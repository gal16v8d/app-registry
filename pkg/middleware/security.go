package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("APP_REGISTRY_ENV") == "prod" && c.Request.URL.Path != "/health" {
			apiToken := c.GetHeader("X-API-Token")
			expectedToken := os.Getenv("APP_REGISTRY_API_TOKEN")

			if apiToken != expectedToken {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		c.Next()
	}
}
