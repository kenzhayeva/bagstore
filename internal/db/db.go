package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "bagstore"
	dbUser := "root"
	dbPass := "root"
	dbPort := "5436"
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация драйвера для миграций
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	// Сначала сбрасываем "грязное" состояние
	if err := m.Force(1); err != nil {
		log.Fatal(err)
	}

	// Затем запускаем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Инициализация GORM после миграций
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}

//
//package db
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//
//	"github.com/joho/godotenv"
//
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//
//	"github.com/golang-migrate/migrate/v4"
//	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
//	_ "github.com/golang-migrate/migrate/v4/source/file"
//)
//
//var DB *gorm.DB
//
//func InitDB() {
//	// Загружаем переменные окружения из .env, если он есть
//	if err := godotenv.Load(); err != nil {
//		log.Println("No .env file found, relying on system environment")
//	}
//
//	// Считываем переменные окружения
//	dbHost := os.Getenv("DB_HOST")
//	dbName := os.Getenv("DB_NAME")
//	dbUser := os.Getenv("DB_USER")
//	dbPass := os.Getenv("DB_PASSWORD")
//	dbPort := os.Getenv("DB_PORT")
//	sslmode := "disable"
//
//	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
//	log.Println("Connecting to:", dbUrl)
//
//	// Подключение через database/sql
//	sqlDB, err := sql.Open("postgres", dbUrl)
//	if err != nil {
//		log.Fatal("Failed to open SQL connection:", err)
//	}
//
//	// Инициализация драйвера для миграций
//	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
//	if err != nil {
//		log.Fatal("Failed to create migration driver:", err)
//	}
//
//	// Определяем абсолютный путь к миграциям
//	cwd, err := os.Getwd()
//	if err != nil {
//		log.Fatal("Failed to get current working directory:", err)
//	}
//	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
//	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
//	log.Println("Using migrations path:", migrationsURL)
//
//	// Запуск миграций
//	m, err := migrate.NewWithDatabaseInstance(migrationsURL, "postgres", driver)
//	if err != nil {
//		log.Fatal("Failed to create migration instance:", err)
//	}
//	if err := m.Up(); err != nil && err.Error() != "no change" {
//		log.Fatal("Migration failed:", err)
//	}
//
//	// GORM
//	gormDB, err := gorm.Open(postgres.New(postgres.Config{
//		Conn: sqlDB,
//	}), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to initialize GORM:", err)
//	}
//	DB = gormDB
//}
