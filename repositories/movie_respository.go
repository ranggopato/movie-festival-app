package repositories

import (
	"movie-festival-app/config"
	"movie-festival-app/models"

	"gorm.io/gorm"
)

func CreateMovie(movie models.Movie) error {
	return config.DB.Create(&movie).Error
}

func UpdateMovie(movieID uint, movieData models.Movie) error {
	return config.DB.Model(&models.Movie{}).Where("id = ?", movieID).Updates(movieData).Error
}

func FindMostViewedMovie() (models.Movie, error) {
	var movie models.Movie
	err := config.DB.Order("views DESC").First(&movie).Error
	return movie, err
}

func FindMostViewedGenres() ([]string, error) {
	var genres []string
	err := config.DB.Raw(`
		SELECT genres 
		FROM movies 
		GROUP BY genres 
		ORDER BY SUM(views) DESC
	`).Scan(&genres).Error
	return genres, err
}

func PaginateMovies(page, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	offset := (page - 1) * limit
	err := config.DB.Offset(offset).Limit(limit).Find(&movies).Error
	return movies, err
}

func SearchMovies(title, description, artists, genres string) ([]models.Movie, error) {
	var movies []models.Movie
	db := config.DB

	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if description != "" {
		db = db.Where("description LIKE ?", "%"+description+"%")
	}
	if artists != "" {
		db = db.Where("artists LIKE ?", "%"+artists+"%")
	}
	if genres != "" {
		db = db.Where("genres LIKE ?", "%"+genres+"%")
	}

	err := db.Find(&movies).Error
	return movies, err
}

func UpdateViewCount(movieID uint) error {
	return config.DB.Model(&models.Movie{}).Where("id = ?", movieID).Update("views", gorm.Expr("views + 1")).Error
}

func LogViewership(viewership models.Viewership) error {
	return config.DB.Create(&viewership).Error
}

func GetMovieByID(movieID uint) (models.Movie, error) {
	var movie models.Movie
	err := config.DB.First(&movie, movieID).Error
	return movie, err
}
