package Controllers

import (
	"github.com/gin-gonic/gin"
	"myApp/Models"
	"myApp/services"
	"net/http"
	"strconv"
)

type MovieController struct {
	service services.MovieServiceInterface
}

func NewMovieController(service services.MovieServiceInterface) *MovieController {
	return &MovieController{service: service}
}

func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.service.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"movies": movies})
}

func (c *MovieController) GetMovieById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	movie, err := c.service.GetMovieBy(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie Models.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	createdMovie, err := c.service.CreateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, createdMovie)
}

func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updatedMovie Models.Movie
	if err := ctx.BindJSON(&updatedMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	movie, err := c.service.UpdateMovie(uint(id), updatedMovie)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.DeleteMovie(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
}
