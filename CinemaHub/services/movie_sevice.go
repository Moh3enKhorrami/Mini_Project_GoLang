package services

import (
	"errors"
	"gorm.io/gorm"
	"myApp/Models"
)

// MovieServiceInterface Interface for movies operation
type MovieServiceInterface interface {
	GetAllMovies() ([]Models.Movie, error)
	GetMovieBy(id uint) (Models.Movie, error)
	CreateMovie(movie Models.Movie) (Models.Movie, error)
	UpdateMovie(id uint, updateMovie Models.Movie) (Models.Movie, error)
	DeleteMovie(id uint) error
}

// MovieService implements the MovieServiceInterface interface
type MovieService struct {
	db *gorm.DB
}

// NewMovieService Constructor function for MovieService
func NewMovieService(db *gorm.DB) MovieServiceInterface {
	return &MovieService{db: db}
}

func (s *MovieService) GetAllMovies() ([]Models.Movie, error) {
	var movies []Models.Movie
	result := s.db.Find(&movies)
	return movies, result.Error
}

func (s *MovieService) GetMovieBy(id uint) (Models.Movie, error) {
	var movie Models.Movie
	result := s.db.Preload("Director").First(&movie, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return movie, errors.New("movie Not Found")
	}
	return movie, result.Error
}

func (s *MovieService) CreateMovie(movie Models.Movie) (Models.Movie, error) {
	result := s.db.Create(&movie)
	return movie, result.Error
}

func (s *MovieService) UpdateMovie(id uint, updateMovie Models.Movie) (Models.Movie, error) {
	var movie Models.Movie
	if err := s.db.First(&movie, id).Error; err != nil {
		return movie, errors.New("movie Not Found")
	}
	movie.Isbn = updateMovie.Isbn
	movie.Title = updateMovie.Title
	movie.Director = updateMovie.Director
	s.db.Save(&movie)
	return movie, nil
}

func (s *MovieService) DeleteMovie(id uint) error {
	var movie Models.Movie
	if err := s.db.Find(&movie, id).Error; err != nil {
		return errors.New("movie Not Found")
	}
	s.db.Delete(&movie)
	return nil
}
