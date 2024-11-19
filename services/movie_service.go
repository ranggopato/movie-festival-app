package services

import (
	"errors"
	"movie-festival-app/models"
	"movie-festival-app/repositories"
	"time"
)

func CreateMovie(movie models.Movie) error {
	return repositories.CreateMovie(movie)
}

func UpdateMovie(movieID uint, movie models.Movie) error {
	return repositories.UpdateMovie(movieID, movie)
}

func GetMostViewedMovie() (models.Movie, error) {
	return repositories.FindMostViewedMovie()
}

func GetMostViewedGenres() ([]string, error) {
	return repositories.FindMostViewedGenres()
}

func ListMovies(page, limit int) ([]models.Movie, error) {
	return repositories.PaginateMovies(page, limit)
}

func SearchMovies(title, description, artists, genres string) ([]models.Movie, error) {
	return repositories.SearchMovies(title, description, artists, genres)
}

func TrackMovieView(movieID uint) error {
	return repositories.UpdateViewCount(movieID)
}

func TrackMovieViewership(userID uint, movieID uint, duration int) error {
	// Periksa apakah film ada
	movie, err := repositories.GetMovieByID(movieID)
	if err != nil {
		return errors.New("film tidak ditemukan")
	}

	// Tambah jumlah views film
	err = repositories.UpdateViewCount(movie.ID)
	if err != nil {
		return err
	}

	// Catat aktivitas menonton
	viewership := models.Viewership{
		UserID:    userID,
		MovieID:   movieID,
		WatchedAt: time.Now(),
		Duration:  duration,
	}
	return repositories.LogViewership(viewership)
}
