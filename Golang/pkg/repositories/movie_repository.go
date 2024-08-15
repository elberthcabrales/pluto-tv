package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/elberthcabrales/movies-api/pkg/cache"
	"github.com/elberthcabrales/movies-api/pkg/models"
)

// MovieRepository defines the interface for movie-related operations
type MovieRepository interface {
	GetMovieByID(id string) (*models.Movie, error)
	GetMovies(page int) (*models.MovieList, error)
	SaveMovie(movie *models.Movie) error
}

type movieRepositoryImpl struct {
	cache     *cache.RedisCache
	apiURL    string
	client    *http.Client
	authToken string
}

// NewMovieRepository creates a new instance of MovieRepository with the provided authentication token and cache.
func NewMovieRepository(authToken string, cache *cache.RedisCache) MovieRepository {
	return &movieRepositoryImpl{
		cache:     cache,
		apiURL:    "https://api.themoviedb.org/3",
		client:    &http.Client{},
		authToken: authToken,
	}
}

func (r *movieRepositoryImpl) GetMovieByID(id string) (*models.Movie, error) {
	// Check if the movie exists in the cache
	log.Printf("Fetching movie with ID %s from cache...", id)
	cachedMovie, err := r.cache.GetValue(id)
	if err == nil && cachedMovie != "" {
		log.Printf("Cache hit for movie ID %s", id)
		var movie models.Movie
		err = json.Unmarshal([]byte(cachedMovie), &movie)
		if err != nil {
			log.Printf("Failed to unmarshal cached movie data for ID %s: %v", id, err)
			return nil, err
		}
		return &movie, nil
	}
	log.Printf("Cache miss for movie ID %s. Fetching from API...", id)
	url := fmt.Sprintf("%s/movie/%s", r.apiURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create HTTP request for movie ID %s: %v", id, err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.authToken))
	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch movie with ID %s from API: %v", id, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned non-200 status code for movie ID %s: %d", id, resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch movie from API, status code: %d", resp.StatusCode)
	}

	var movie models.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		log.Printf("Failed to decode API response for movie ID %s: %v", id, err)
		return nil, err
	}

	movieJSON, err := json.Marshal(movie)
	if err != nil {
		log.Printf("Failed to marshal movie data for caching: %v", err)
		return nil, err
	}
	err = r.cache.SetValue(id, string(movieJSON))
	if err != nil {
		log.Printf("Failed to cache movie data for ID %s: %v", id, err)
		return nil, err
	}

	log.Printf("Movie with ID %s fetched and cached successfully", id)
	return &movie, nil
}

func (r *movieRepositoryImpl) GetMovies(page int) (*models.MovieList, error) {
	url := fmt.Sprintf("%s/discover/movie?page=%d", r.apiURL, page)
	log.Printf("Fetching movies from URL: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create HTTP request for movies: %v", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.authToken))
	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Failed to fetch movies from API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("API returned non-200 status code for movies: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch movies from API, status code: %d", resp.StatusCode)
	}

	var response models.MovieList
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Printf("Failed to decode API response for movies: %v", err)
		return nil, err
	}

	log.Printf("Successfully fetched movies from API: %v", response)
	return &response, nil
}

// Note: if the movie exists in the cache, it will be overwritten with the new data
func (r *movieRepositoryImpl) SaveMovie(movie *models.Movie) error {
	// Serialize the movie
	movieJSON, err := json.Marshal(movie)
	if err != nil {
		log.Printf("Failed to marshal movie data for ID %d: %v", movie.ID, err)
		return err
	}
	id := strconv.Itoa(movie.ID)
	// Save to cache
	err = r.cache.SetValue(id, string(movieJSON))
	if err != nil {
		log.Printf("Failed to save movie with ID %d to cache: %v", movie.ID, err)
		return err
	}

	log.Printf("Movie with ID %d saved to cache successfully", movie.ID)
	return nil
}
