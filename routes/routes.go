package routes

import (
	"urlShorter/controllers"
	"urlShorter/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	urlService, _ := services.NewURLService(1)
	urlController := controllers.NewURLController(urlService)

	r.POST("/shorten", urlController.ShortenURL)
	r.GET("/:short_url", urlController.RedirectURL)
	r.GET("/qrcode/:short_url", urlController.GetQRCode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	return r
}
