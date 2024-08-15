package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/elberthcabrales/movies-api/pkg/models"
)

type MockMovieService struct {
	mock.Mock
}

func (m *MockMovieService) GetMovieByID(id string) (*models.Movie, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Movie), args.Error(1)
}

func (m *MockMovieService) GetMovies(page int) (*models.MovieList, error) {
	args := m.Called(page)
	return args.Get(0).(*models.MovieList), args.Error(1)
}

func (m *MockMovieService) SaveMovie(movie *models.Movie) error {
	args := m.Called(movie)
	return args.Error(0)
}

func TestGetMovieByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockMovieService)
	router := NewMovieRouter(mockService).SetupRouter()

	expectedMovie := &models.Movie{
		ID:    1,
		Title: "Test Movie",
	}

	mockService.On("GetMovieByID", "1").Return(expectedMovie, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualMovie models.Movie
	err := json.Unmarshal(w.Body.Bytes(), &actualMovie)
	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, &actualMovie)
}

func TestGetMovieByID_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockMovieService)
	router := NewMovieRouter(mockService).SetupRouter()

	mockService.On("GetMovieByID", "1").Return((*models.Movie)(nil), errors.New("service error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "service error"}`, w.Body.String())
}

func TestGetMovies(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockMovieService)
	router := NewMovieRouter(mockService).SetupRouter()

	expectedMovies := &models.MovieList{
		Results: []models.Movie{
			{ID: 1, Title: "Movie 1"},
			{ID: 2, Title: "Movie 2"},
			{ID: 3, Title: "Movie 3"},
		},
		Page: 1,
	}

	mockService.On("GetMovies", 1).Return(expectedMovies, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies?page=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualMovies models.MovieList
	err := json.Unmarshal(w.Body.Bytes(), &actualMovies)
	assert.NoError(t, err)
	assert.Equal(t, *expectedMovies, actualMovies)
}

func TestGetMovies_InvalidPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewMovieRouter(nil).SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/movies?page=invalid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "Invalid page value"}`, w.Body.String())
}

func TestSaveMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockMovieService)
	router := NewMovieRouter(mockService).SetupRouter()

	newMovie := &models.Movie{
		ID:    1,
		Title: "New Movie",
	}

	mockService.On("SaveMovie", newMovie).Return(nil)

	jsonMovie, err := json.Marshal(newMovie)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(jsonMovie))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSaveMovie_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockMovieService)
	router := NewMovieRouter(mockService).SetupRouter()

	movieToSave := &models.Movie{
		ID:    1,
		Title: "Test Movie",
	}

	mockService.On("SaveMovie", movieToSave).Return(errors.New("service error"))

	movieJSON, _ := json.Marshal(movieToSave)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(movieJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error": "service error"}`, w.Body.String())

	mockService.AssertExpectations(t)
}
