//package main
//
//import (
//	"bagstore/internal/models"
//	"bagstore/internal/routes"
//	"github.com/gin-gonic/gin"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//	"log"
//)
//
//func main() {
//	dsn := "postgres://root:mypassword@localhost:5436/root?sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Error connecting to the database:", err)
//	}
//
//	err = db.AutoMigrate(&models.Bag{})
//	if err != nil {
//		log.Fatal("Error on migrating to the DB:", err)
//	}
//
//	r := gin.Default()
//	routes.SetupRoutes(r, db)
//	r.Run(":8081")
//}

package main

import (
	"bagstore/internal/models"
	"bagstore/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Строка подключения с явным указанием параметров
	dsn := "host=localhost port=5436 user=root password=root dbname=bagstore sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Миграция базы данных
	err = db.AutoMigrate(&models.Bag{})
	if err != nil {
		log.Fatal("Error on migrating to the DB:", err)
	}

	// Инициализация маршрутов
	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8081")
}
