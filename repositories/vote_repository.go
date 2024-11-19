package repositories

import (
	"movie-festival-app/config"
	"movie-festival-app/models"
)

// AddVote menambahkan vote untuk sebuah film
func AddVote(vote models.Vote) error {
	return config.DB.Create(&vote).Error
}

// RemoveVote menghapus vote untuk sebuah film
func RemoveVote(userID, movieID uint) error {
	return config.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Vote{}).Error
}

// HasUserVoted memeriksa apakah pengguna sudah mem-vote film tertentu
func HasUserVoted(userID, movieID uint) (bool, error) {
	var vote models.Vote
	err := config.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).First(&vote).Error
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetUserVotedMovies mendapatkan semua film yang telah di-vote oleh pengguna
func GetUserVotedMovies(userID uint) ([]models.Movie, error) {
	var movies []models.Movie
	err := config.DB.Table("movies").
		Select("movies.*").
		Joins("join votes on votes.movie_id = movies.id").
		Where("votes.user_id = ?", userID).
		Find(&movies).Error
	return movies, err
}

// GetMostVotedMovies mendapatkan film yang paling banyak di-vote
func GetMostVotedMovies(limit int) ([]models.MovieWithVotes, error) {
	var movies []models.MovieWithVotes
	err := config.DB.Table("movies").
		Select("movies.*, COUNT(votes.movie_id) as vote_count").
		Joins("left join votes on votes.movie_id = movies.id").
		Group("movies.id").
		Order("vote_count DESC").
		Limit(limit).
		Scan(&movies).Error
	return movies, err
}
