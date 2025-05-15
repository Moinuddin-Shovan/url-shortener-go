package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moinuddinshovan/url-shortener-go/internal/handlers"
	"github.com/moinuddinshovan/url-shortener-go/internal/models"
	"github.com/moinuddinshovan/url-shortener-go/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.URL{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize services
	urlService := services.NewURLService(db)

	// Initialize handlers
	urlHandler := handlers.NewURLHandler(urlService)

	// Initialize router
	router := gin.Default()

	// Register routes
	urlHandler.RegisterRoutes(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
