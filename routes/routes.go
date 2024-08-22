package routes

import (
	urlHandler "github.com/BAXF/shortener/api/handlers/url"
	middlewares "github.com/BAXF/shortener/api/middleware"

	"github.com/BAXF/shortener/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, redisClient *models.RedisClient) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.GoogleAuthMiddleware())

	r.POST("/shorten", urlHandler.CreateURL)
	r.GET("/:shortURL", urlHandler.GetOriginalURL)

	return r
}
