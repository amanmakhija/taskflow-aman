package main

import (
	"log"
	"net/http"
	"taskflow/internal/auth"
	"taskflow/internal/config"
	"taskflow/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db.Connect(cfg.DatabaseURL)
	r := gin.Default()

	authService := &auth.Service{JWTSecret: cfg.JWTSecret}
	authHandler := &auth.Handler{Service: authService}

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}
