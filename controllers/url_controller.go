package controllers

import (
	"net/http"
	"urlShorter/services"

	"github.com/gin-gonic/gin"
)

type URLController struct {
	Service *services.URLService
}

func NewURLController(service *services.URLService) *URLController {
	return &URLController{Service: service}
}

func (uc *URLController) ShortenURL(c *gin.Context) {
	var requestBody struct {
		LongURL string `json:"long_url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	shortURL, err := uc.Service.ShortenURL(requestBody.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating the link"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func (uc *URLController) RedirectURL(c *gin.Context) {
	shortURL := c.Param("short_url")

	longURL, err := uc.Service.GetOriginalURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}

func (uc *URLController) GetQRCode(c *gin.Context) {
	shortURL := c.Param("short_url")

	qrCode, err := uc.Service.GenerateQRCode(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	c.Data(http.StatusOK, "image/png", qrCode)
}
