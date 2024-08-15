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

// NewMovieRouter creates a new MovieRouter.
func NewMovieRouter(movieService services.MovieService) *MovieRouter {
	return &MovieRouter{movieService: movieService}
}

// SetupRouter configures the routes for the movie router.
func (r *MovieRouter) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/movies/:id", r.getMovieByID)
	router.GET("/movies", r.getMovies)
	router.POST("/movies", r.saveMovie)

	return router
}

func (r *MovieRouter) getMovieByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := r.movieService.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

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
