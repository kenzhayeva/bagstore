package routes

import (
	"bagstore/internal/delivery"
	"bagstore/internal/repository"
	"bagstore/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Инициализация репозитория
	bagRepo := repository.NewBagRepository(db)

	// Инициализация сервиса
	bagService := services.NewBagService(bagRepo)

	// Инициализация обработчика
	bagHandler := delivery.NewBagHandler(bagService)

	// Роуты
	bags := r.Group("/api/v1/bags")
	{
		bags.GET("/", bagHandler.GetAllBags)
		bags.GET("/:id", bagHandler.GetBag)
		bags.POST("/", bagHandler.CreateBag)
		bags.PUT("/:id", bagHandler.UpdateBag)
		bags.DELETE("/:id", bagHandler.DeleteBag)
	}
}
