package routes

import (
	"github.com/BAXF/shortener/api/handlers"
	"github.com/BAXF/shortener/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, redisClient *models.RedisClient) *gin.Engine {
	r := gin.Default()

	urlHandler := handlers.URLHanlder{DB: db, Redis: redisClient}
	r.POST("/shorten", urlHandler.CreateURL)
	r.GET("/:shortURL", urlHandler.GetOriginalURL)

	return r
}
