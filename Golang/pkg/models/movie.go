package models

// Genre represents a movie genre with an ID and a name.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProductionCompany represents a company involved in producing a movie.
type ProductionCompany struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// ProductionCountry represents a country where a movie was produced.
type ProductionCountry struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

// SpokenLanguage represents a language spoken in a movie.
type SpokenLanguage struct {
	EnglishName string `json:"english_name"`
	Iso639_1    string `json:"iso_639_1"`
	Name        string `json:"name"`
}

// Movie represents the details of a movie.
type Movie struct {
	ID                  int                 `json:"id"`
	Title               string              `json:"title"`
	OriginalTitle       string              `json:"original_title"`
	Overview            string              `json:"overview"`
	ReleaseDate         string              `json:"release_date"`
	Runtime             int                 `json:"runtime"`
	VoteAverage         float64             `json:"vote_average"`
	VoteCount           int                 `json:"vote_count"`
	PosterPath          string              `json:"poster_path"`
	BackdropPath        string              `json:"backdrop_path"`
	Genres              []Genre             `json:"genres"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	SpokenLanguages     []SpokenLanguage    `json:"spoken_languages"`
	Budget              int                 `json:"budget"`
	Revenue             int                 `json:"revenue"`
	Homepage            string              `json:"homepage"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Video               bool                `json:"video"`
	Adult               bool                `json:"adult"`
}

// MovieList represents a paginated list of movies.
type MovieList struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Message string `json:"message"`
}
