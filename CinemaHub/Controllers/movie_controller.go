package Controllers

import (
	"github.com/gin-gonic/gin"
	"myApp/models"
	"myApp/services"
	"net/http"
	"strconv"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// MovieController handles movie related requests
type MovieController struct {
	service services.MovieServiceInterface
}

// NewMovieController creates a new MovieControllerss
func NewMovieController(service services.MovieServiceInterface) *MovieController {
	return &MovieController{service}
}

// GetAllMovies godoc
// @Summary Get all movies
// @Description Get all movies
// @Tags movies
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Movie
// @Router /movies [get]
func (c *MovieController) GetAllMovies(ctx *gin.Context) {
	movies, err := c.service.GetAllMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

// GetMovieByID godoc
// @Summary Get a movie by ID
// @Description Get a movie by ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [get]
func (c *MovieController) GetMovieById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	movie, err := c.service.GetMovieBy(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Message: "not Find"})
	}
	ctx.JSON(http.StatusOK, movie)
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Create a new movie
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body models.Movie true "Movie to create"
// @Success 201 {object} models.Movie
// @Failure 400 {object} ErrorResponse
// @Router /movies [post]
func (c *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "Bad Request"})
		return
	}
	createdMovie, err := c.service.CreateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal Server Error"})
	}
	ctx.JSON(http.StatusCreated, createdMovie)
}

// UpdateMovie godoc
// @Summary Update a movie
// @Description Update a movie
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Param movie body models.Movie true "Updated movie"
// @Success 200 {object} models.Movie
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [put]
func (c *MovieController) UpdateMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updatedMovie models.Movie
	if err := ctx.BindJSON(&updatedMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: "Bad Request"})
		return
	}
	movie, err := c.service.UpdateMovie(uint(id), updatedMovie)
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Message: "Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete a movie
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 404 {object} ErrorResponse
// @Router /movies/{id} [delete]
func (c *MovieController) DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.DeleteMovie(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse{Message: "Not Found"})
		return
	}
	ctx.JSON(http.StatusOK, ErrorResponse{Message: "Status Error"})
}
