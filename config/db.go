package config

import (
	"errors"
	"fmt"
	"github.com/dnlgrwd87/blog-api/models"
	"github.com/go-chi/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitializeDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return &gorm.DB{}, errors.New(fmt.Sprintf("couldn't connect to the database: %v", err.Error()))
	}

	db.LogMode(true)

	fmt.Println("successfully connected to the database")
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Post{})
}

func AddContstraints(db *gorm.DB) {
	db.Model(&models.Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}

func GetCors() *cors.Cors {
	corsOptions := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		})
	return corsOptions
}
