package models

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenre_MarshalUnmarshalJSON(t *testing.T) {
	genre := Genre{
		ID:   1,
		Name: "Action",
	}

	data, err := json.Marshal(genre)
	assert.NoError(t, err)

	var unmarshaledGenre Genre
	err = json.Unmarshal(data, &unmarshaledGenre)
	assert.NoError(t, err)
	assert.Equal(t, genre, unmarshaledGenre)
}

func TestProductionCompany_MarshalUnmarshalJSON(t *testing.T) {
	company := ProductionCompany{
		ID:            1,
		LogoPath:      "/path/to/logo.png",
		Name:          "Company Name",
		OriginCountry: "US",
	}

	data, err := json.Marshal(company)
	assert.NoError(t, err)

	var unmarshaledCompany ProductionCompany
	err = json.Unmarshal(data, &unmarshaledCompany)
	assert.NoError(t, err)
	assert.Equal(t, company, unmarshaledCompany)
}

func TestProductionCountry_MarshalUnmarshalJSON(t *testing.T) {
	country := ProductionCountry{
		Iso3166_1: "US",
		Name:      "United States of America",
	}

	data, err := json.Marshal(country)
	assert.NoError(t, err)

	var unmarshaledCountry ProductionCountry
	err = json.Unmarshal(data, &unmarshaledCountry)
	assert.NoError(t, err)
	assert.Equal(t, country, unmarshaledCountry)
}

func TestSpokenLanguage_MarshalUnmarshalJSON(t *testing.T) {
	language := SpokenLanguage{
		EnglishName: "English",
		Iso639_1:    "en",
		Name:        "English",
	}

	data, err := json.Marshal(language)
	assert.NoError(t, err)

	var unmarshaledLanguage SpokenLanguage
	err = json.Unmarshal(data, &unmarshaledLanguage)
	assert.NoError(t, err)
	assert.Equal(t, language, unmarshaledLanguage)
}

func TestMovie_MarshalUnmarshalJSON(t *testing.T) {
	movie := Movie{
		ID:            1,
		Title:         "Test Movie",
		OriginalTitle: "Test Original Title",
		Overview:      "Test overview of the movie.",
		ReleaseDate:   "2024-06-05",
		Runtime:       120,
		VoteAverage:   8.7,
		VoteCount:     1500,
		PosterPath:    "/path/to/poster.png",
		BackdropPath:  "/path/to/backdrop.png",
		Genres: []Genre{
			{ID: 1, Name: "Action"},
			{ID: 2, Name: "Comedy"},
		},
		ProductionCompanies: []ProductionCompany{
			{ID: 1, LogoPath: "/path/to/logo.png", Name: "Company Name", OriginCountry: "US"},
		},
		ProductionCountries: []ProductionCountry{
			{Iso3166_1: "US", Name: "United States of America"},
		},
		SpokenLanguages: []SpokenLanguage{
			{EnglishName: "English", Iso639_1: "en", Name: "English"},
		},
		Budget:   100000000,
		Revenue:  200000000,
		Homepage: "http://example.com",
		Status:   "Released",
		Tagline:  "This is a test tagline.",
		Video:    false,
		Adult:    false,
	}

	data, err := json.Marshal(movie)
	assert.NoError(t, err)

	var unmarshaledMovie Movie
	err = json.Unmarshal(data, &unmarshaledMovie)
	assert.NoError(t, err)
	assert.Equal(t, movie, unmarshaledMovie)
}

func TestMovieList_MarshalUnmarshalJSON(t *testing.T) {
	movieList := MovieList{
		Page: 1,
		Results: []Movie{
			{ID: 1, Title: "Test Movie 1"},
			{ID: 2, Title: "Test Movie 2"},
		},
	}

	data, err := json.Marshal(movieList)
	assert.NoError(t, err)

	var unmarshaledMovieList MovieList
	err = json.Unmarshal(data, &unmarshaledMovieList)
	assert.NoError(t, err)
	assert.Equal(t, movieList, unmarshaledMovieList)
}

func TestErrorResponse_MarshalUnmarshalJSON(t *testing.T) {
	errorResponse := ErrorResponse{
		Error: "This is an error message.",
	}

	data, err := json.Marshal(errorResponse)
	assert.NoError(t, err)

	var unmarshaledErrorResponse ErrorResponse
	err = json.Unmarshal(data, &unmarshaledErrorResponse)
	assert.NoError(t, err)
	assert.Equal(t, errorResponse, unmarshaledErrorResponse)
}

func TestSuccessResponse_MarshalUnmarshalJSON(t *testing.T) {
	successResponse := SuccessResponse{
		Message: "This is a success message.",
	}

	data, err := json.Marshal(successResponse)
	assert.NoError(t, err)

	var unmarshaledSuccessResponse SuccessResponse
	err = json.Unmarshal(data, &unmarshaledSuccessResponse)
	assert.NoError(t, err)
	assert.Equal(t, successResponse, unmarshaledSuccessResponse)
}

func TestMovieJSONMarshalling(t *testing.T) {
	movieData, err := loadMockData("movie_573435.json")
	if err != nil {
		t.Fatalf("Failed to load mock data: %v", err)
	}

	var movie Movie
	err = json.Unmarshal(movieData, &movie)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	marshalledData, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Failed to marshal movie to JSON: %v", err)
	}

	var movieCopy Movie
	err = json.Unmarshal(marshalledData, &movieCopy)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	assertEqual(t, movie.ID, movieCopy.ID, "movie ID")
	assertEqual(t, movie.Title, movieCopy.Title, "movie title")
}

func TestMoviesJSONMarshalling(t *testing.T) {
	moviesData, err := loadMockData("movies_list.json")
	if err != nil {
		t.Fatalf("Failed to load mock data: %v", err)
	}

	var movies MovieList
	err = json.Unmarshal(moviesData, &movies)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	marshalledData, err := json.Marshal(movies)
	if err != nil {
		t.Fatalf("Failed to marshal movies to JSON: %v", err)
	}

	var moviesCopy MovieList
	err = json.Unmarshal(marshalledData, &moviesCopy)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	assertEqual(t, movies.Page, moviesCopy.Page, "movies page")
	assertEqual(t, len(movies.Results), len(moviesCopy.Results), "movies results length")
}

func assertEqual(t *testing.T, got, want interface{}, name string) {
	if got != want {
		t.Errorf("Expected %s to be %v, got %v", name, want, got)
	}
}

func loadMockData(filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join("../../testdata/mocks", filename))
}
