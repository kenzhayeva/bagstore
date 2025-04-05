package routes

import (
	"bagstore/internal/delivery"
	"bagstore/internal/repository"
	"bagstore/internal/services"

	"github.com/gin-gonic/gin"
)


	bagService := services.NewBagService(bagRepo)
	bagHandler := delivery.NewBagHandler(bagService)

	{
		bags.GET("/", bagHandler.GetAllBags)
		bags.GET("/:id", bagHandler.GetBag)
		bags.POST("/", bagHandler.CreateBag)
		bags.PUT("/:id", bagHandler.UpdateBag)
		bags.DELETE("/:id", bagHandler.DeleteBag)
	}
}
