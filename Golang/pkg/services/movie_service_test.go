package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/elberthcabrales/movies-api/pkg/models"
)

type MockMovieRepository struct {
	mock.Mock
}

func (m *MockMovieRepository) GetMovieByID(id string) (*models.Movie, error) {
	args := m.Called(id)
	if movie, ok := args.Get(0).(*models.Movie); ok {
		return movie, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMovieRepository) GetMovies(page int) (*models.MovieList, error) {
	args := m.Called(page)
	if movies, ok := args.Get(0).(*models.MovieList); ok {
		return movies, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockMovieRepository) SaveMovie(movie *models.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func TestMovieService_GetMovieByID(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	service := NewMovieService(mockRepo)

	expectedMovie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	mockRepo.On("GetMovieByID", "573435").Return(expectedMovie, nil)

	movie, err := service.GetMovieByID("573435")

	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)

	mockRepo.AssertExpectations(t)
}

func TestMovieService_GetMovies(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	service := NewMovieService(mockRepo)

	expectedMovies := &models.MovieList{
		Results: []models.Movie{
			{ID: 533535, Title: "Dead pool & Wolverine"},
			{ID: 573435, Title: "Bad Boys: Ride or Die"},
		},
		Page: 1,
	}

	mockRepo.On("GetMovies", 1).Return(expectedMovies, nil)

	movies, err := service.GetMovies(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, movies)

	mockRepo.AssertExpectations(t)
}

func TestMovieService_SaveMovie(t *testing.T) {
	mockRepo := new(MockMovieRepository)
	service := NewMovieService(mockRepo)

	movie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	mockRepo.On("SaveMovie", movie).Return(nil)

	err := service.SaveMovie(movie)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
