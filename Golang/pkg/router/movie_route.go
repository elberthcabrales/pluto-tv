package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/elberthcabrales/movies-api/pkg/models"
	"github.com/elberthcabrales/movies-api/pkg/services"
)

// MovieRouter handles the routing of movie-related HTTP requests.
type MovieRouter struct {
	movieService services.MovieService
}

// NewMovieRouter creates a new MovieRouter
func NewMovieRouter(movieService services.MovieService) *MovieRouter {
	return &MovieRouter{movieService: movieService}
}

// SetupRouter sets up the routes for the MovieRouter
// @title Movie API
// @version 1.0
// @description This is a sample server for managing movies.
// @host localhost:8080
// @BasePath /
func (r *MovieRouter) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/movies/:id", r.getMovieByID)
	router.GET("/movies", r.getMovies)
	router.POST("/movies", r.saveMovie)

	return router
}

// getMovieByID godoc
// @Summary Get a movie by ID
// @Description Get a movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path string true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 500 {object} models.ErrorResponse
// @Router /movies/{id} [get]
func (r *MovieRouter) getMovieByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := r.movieService.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// getMovies godoc
// @Summary Get movies
// @Description Get a list of movies
// @Tags movies
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Success 200 {array} models.Movie
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /movies [get]
func (r *MovieRouter) getMovies(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid page value"})
		return
	}
	movies, err := r.movieService.GetMovies(pageInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// saveMovie godoc
// @Summary Save a movie
// @Description Save a movie in the cache
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body models.Movie true "Movie to save"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /movies [post]
func (r *MovieRouter) saveMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := r.movieService.SaveMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Movie saved successfully"})
}
