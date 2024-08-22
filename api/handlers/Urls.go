package handlers

import (
	"github.com/BAXF/shortener/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

type URLHanlder struct {
	DB    *gorm.DB
	Redis *models.RedisClient
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateShortURL() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func (h *URLHanlder) CreateURL(c *gin.Context) {
	var input struct {
		Original string `json:"original" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := generateShortURL()

	url := models.URL{
		Original: input.Original,
		ShortURL: shortURL,
		UserID:   c.GetString("userID"),
	}
	h.DB.Create(&url)

	c.JSON(http.StatusOK, gin.H{
		"shortened": shortURL,
	})
}

func (h *URLHanlder) GetOriginalURL(c *gin.Context) {
	shortURL := c.Param("shortURL")

	var url models.URL
	if err := h.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url.Original)
}
