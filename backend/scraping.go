package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type URLToScrape struct {
	gorm.Model
	URL string `json:"url"`
}

type ScrapedURL struct {
	gorm.Model
	URL         string `json:"url"`
	LinksTo     uint   `json:"linksto"`
	LinkedTo    uint   `json:"linkedto"`
	Title       string `json:"pagetitle"`
	HTMLContent string `json:"html"`
}

func AddURLToScrape(c *gin.Context) {
	var uts URLToScrape
	c.ShouldBindJSON(&uts)

	err := db.Create(&uts).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"message": "URL successfully added",
	})
}

func GetURLToScrape(c *gin.Context) {
	var uts URLToScrape
	err := db.First(&uts).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = db.Where("ID = ?", uts.ID).Delete(&uts).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"url": uts.URL,
	})
}

func AddScrapedURL(c *gin.Context) {
	var us ScrapedURL
	c.ShouldBindJSON(&us)

	err := db.Create(&us).Error

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{
		"message": "URL successfully added",
	})
}
