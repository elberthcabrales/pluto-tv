package services

import (
	"github.com/elberthcabrales/movies-api/pkg/models"
	"github.com/elberthcabrales/movies-api/pkg/repositories"
)

// MovieService defines the interface for movie-related operations
type MovieService interface {
	GetMovieByID(id string) (*models.Movie, error)
	GetMovies(page int) (*models.MovieList, error)
	SaveMovie(movie *models.Movie) error
}

type movieService struct {
	repo repositories.MovieRepository
}

// NewMovieService creates a new instance of MovieService
func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

// GetMovieByID retrieves a movie by its ID, either from the cache or the API
func (s *movieService) GetMovieByID(id string) (*models.Movie, error) {
	return s.repo.GetMovieByID(id)
}

// GetMovies retrieves a list of movies from the API
func (s *movieService) GetMovies(page int) (*models.MovieList, error) {
	return s.repo.GetMovies(page)
}

// SaveMovie saves a movie in the cache
func (s *movieService) SaveMovie(movie *models.Movie) error {
	return s.repo.SaveMovie(movie)
}
