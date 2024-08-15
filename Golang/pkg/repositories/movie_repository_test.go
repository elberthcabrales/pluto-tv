package repositories

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"

	"github.com/elberthcabrales/movies-api/pkg/cache"
	"github.com/elberthcabrales/movies-api/pkg/models"
)

func TestGetMovieByID_CacheHit(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	expectedMovie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	movieJSON, _ := json.Marshal(expectedMovie)

	mock.ExpectGet("573435").SetVal(string(movieJSON))

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	movie, err := repo.GetMovieByID("573435")

	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieByID_APIHit(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	expectedMovie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	mock.ExpectGet("573435").RedisNil()

	apiResponse, _ := json.Marshal(expectedMovie)
	defer gock.Off()
	gock.New("https://api.themoviedb.org/3").
		Get("/movie/573435").
		MatchHeader("Authorization", "Bearer dummy-auth-token").
		Reply(200).
		JSON(apiResponse)

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	mock.ExpectSet("573435", string(apiResponse), 0).SetVal("OK")

	movie, err := repo.GetMovieByID("573435")

	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, movie)

	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestGetMovieByID_APINotFound(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	mock.ExpectGet("573435").RedisNil()

	gock.New("https://api.themoviedb.org/3").
		Get("/movie/573435").
		MatchHeader("Authorization", "Bearer dummy-auth-token").
		Reply(404)

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	movie, err := repo.GetMovieByID("573435")

	assert.Error(t, err)
	assert.Nil(t, movie)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovieByID_APIUnavailable(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	mock.ExpectGet("573435").RedisNil()

	gock.New("https://api.themoviedb.org/3").
		Get("/movie/573435").
		MatchHeader("Authorization", "Bearer dummy-auth-token").
		ReplyError(fmt.Errorf("API unavailable"))

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	movie, err := repo.GetMovieByID("573435")

	assert.Error(t, err)
	assert.Nil(t, movie)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMovies(t *testing.T) {
	expectedMovies := []models.Movie{
		{ID: 533535, Title: "Deadpool & Wolverine"},
		{ID: 573435, Title: "Bad Boys: Ride or Die"},
	}

	apiResponse, _ := json.Marshal(models.MovieList{
		Page:    1,
		Results: expectedMovies,
	})
	defer gock.Off()
	gock.New("https://api.themoviedb.org/3").
		Get("/discover/movie").
		MatchHeader("Authorization", "Bearer dummy-auth-token").
		MatchParam("page", "1").
		Reply(200).
		JSON(apiResponse)

	repo := NewMovieRepository("dummy-auth-token", nil)

	movieList, err := repo.GetMovies(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, movieList.Page)
	assert.Equal(t, expectedMovies, movieList.Results)
}

func TestSaveMovie(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	movie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	movieJSON, _ := json.Marshal(movie)

	mock.ExpectSet("573435", string(movieJSON), 0).SetVal("OK")

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	err := repo.SaveMovie(movie)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveMovie_Failure(t *testing.T) {
	db, mock := redismock.NewClientMock()
	redisCache := cache.NewRedisCache(db, nil)

	movie := &models.Movie{
		ID:    573435,
		Title: "Bad Boys: Ride or Die",
	}

	movieJSON, _ := json.Marshal(movie)

	mock.ExpectSet("573435", string(movieJSON), 0).SetErr(fmt.Errorf("failed to save movie"))

	repo := NewMovieRepository("dummy-auth-token", redisCache)

	err := repo.SaveMovie(movie)

	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
