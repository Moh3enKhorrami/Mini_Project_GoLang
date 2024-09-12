package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"myApp/Controllers"
	_ "myApp/docs"
	"myApp/models"
	"myApp/services"
)

// @title CinemaHub API
// @version 1.0
// @description This is a sample server for CinemaHub.
// @host localhost:8080
// @BasePath /
func main() {
	dsn := "postgresql://postgres:123456@localhost:5432/TestDB?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// Automatic migration of models
	_ = db.AutoMigrate(&models.Movie{}, &models.Director{})

	// Create service and video controller
	movieService := services.NewMovieService(db)
	movieController := Controllers.NewMovieController(movieService)

	// Setup Gin router
	r := gin.Default()

	// Register Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set API routes
	r.GET("/movies", movieController.GetAllMovies)
	r.GET("/movies/:id", movieController.GetMovieById)
	r.POST("/movies", movieController.CreateMovie)
	r.PUT("/movies/:id", movieController.UpdateMovie)
	r.DELETE("/movies/:id", movieController.DeleteMovie)

	r.Run("localhost:8080")
}
