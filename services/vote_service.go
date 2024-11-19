package services

import (
	"errors"
	"movie-festival-app/models"
	"movie-festival-app/repositories"
)

// VoteMovie memberikan suara untuk sebuah film
func VoteMovie(userID, movieID uint) error {

	// Periksa apakah film ada
	_, err := repositories.GetMovieByID(movieID)
	if err != nil {
		return errors.New("film tidak ditemukan")
	}

	// Cek apakah user sudah memberikan vote
	hasVoted, err := repositories.HasUserVoted(userID, movieID)
	if err != nil {
		return err
	}

	if hasVoted {
		return errors.New("user sudah memberikan vote untuk film ini")
	}

	// Tambahkan vote
	vote := models.Vote{
		UserID:  userID,
		MovieID: movieID,
	}
	return repositories.AddVote(vote)
}

// UnvoteMovie unvote dari sebuah film
func UnvoteMovie(userID, movieID uint) error {

	_, err := repositories.GetMovieByID(movieID)
	if err != nil {
		return errors.New("film tidak ditemukan")
	}

	return repositories.RemoveVote(userID, movieID)
}

// ListUserVotedMovies mendapatkan daftar film yang di-vote oleh user
func ListUserVotedMovies(userID uint) ([]models.Movie, error) {
	return repositories.GetUserVotedMovies(userID)
}

// GetMostVotedMovies mendapatkan daftar film yang paling banyak di-vote
func GetMostVotedMovies(limit int) ([]models.MovieWithVotes, error) {
	return repositories.GetMostVotedMovies(limit)
}
