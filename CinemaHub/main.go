package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"myApp/Controllers"
	"myApp/Models"
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
	_ = db.AutoMigrate(&Models.Movie{}, &Models.Director{})

	// Create service and video controller
	movieService := services.NewMovieService(db)
	movieController := Controllers.NewMovieController(movieService)

	// Set API routes
	r := gin.Default()
	r.GET("/movie", movieController.GetAllMovies)
	r.GET("/movie/:id", movieController.GetMovieById)
	r.POST("/movie", movieController.CreateMovie)
	r.PUT("/movie/:id", movieController.UpdateMovie)
	r.DELETE("/movie/:id", movieController.DeleteMovie)

	r.Run("localhost:8080")
}
