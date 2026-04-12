package main

import (
	"log"
	"net/http"
	"taskflow/internal/config"
	"taskflow/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db.Connect(cfg.DatabaseURL)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}
