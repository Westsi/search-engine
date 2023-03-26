package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db = _db

	// Migrate the schema
	db.AutoMigrate(&URLToScrape{}, &ScrapedURL{})
}
