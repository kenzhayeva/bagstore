package main

import (
	"bagstore/internal/db"
	"bagstore/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	r := gin.Default()

	db.InitDB()
	routes.SetupRoutes(r, db.DB)
	routes.RegisterAdminRoutes(r, db.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
