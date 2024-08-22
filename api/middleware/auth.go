package middlewares

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v2"
	"net/http"
)

func GoogleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		oauth2Service, err := oauth2.NewService(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OAuth2 service"})
			c.Abort()
			return
		}

		tokenInfo, err := oauth2Service.Tokeninfo().AccessToken(token).Do()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", tokenInfo.UserId)
		c.Next()
	}
}
