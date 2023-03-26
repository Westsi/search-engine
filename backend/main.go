package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	fmt.Println("Error loading .env file - does it exist?")
	// 	return
	// }

	initDB()
	r := gin.Default()

	routerConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(routerConfig))
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "pong",
		})
	})

	scraping := r.Group("/scraping")
	scraping.GET("/nexturl", GetURLToScrape)
	scraping.POST("/addscrapedurl", AddScrapedURL)
	scraping.POST("/addurltoscrape", AddURLToScrape)

	r.Run("localhost:8080")
}
