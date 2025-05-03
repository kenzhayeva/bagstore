package routes

import (
	"bagstore/internal/admin"
	"github.com/gin-gonic/gin"

	"bagstore/internal/authorization"
	"bagstore/internal/delivery"
	"bagstore/internal/middleware"
	"bagstore/internal/repository"
	"bagstore/internal/services"

	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Публичные роуты
	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", authorization.Login)
		authRoutes.POST("/register", authorization.Register)

		// Защищённые роуты
		protected := r.Group("api/v1")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/me", authorization.Me)

			// Бизнес-логика сумок (только авторизованные пользователи)
			bagRepo := repository.NewBagRepository(db)
			bagService := services.NewBagService(bagRepo)
			bagHandler := delivery.NewBagHandler(bagService)

			bags := protected.Group("/bags")
			{
				bags.GET("/", bagHandler.GetAllBags)
				bags.GET("/:id", bagHandler.GetBag)
				bags.POST("/", bagHandler.CreateBag)
				bags.PUT("/:id", bagHandler.UpdateBag)
				bags.DELETE("/:id", bagHandler.DeleteBag)
			}
		}
	}
}

func RegisterAdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminHandler := &admin.AdminHandler{DB: db}

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AuthRequired(), middleware.AdminOnly()) // 💡 подключаем middleware
	{
		adminGroup.DELETE("/users", adminHandler.DeleteAllUsers)
		adminGroup.DELETE("/users/:id", adminHandler.DeleteUserByID)
	}
}
