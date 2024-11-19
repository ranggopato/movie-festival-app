package repositories

import (
	"movie-festival-app/config"
	"movie-festival-app/models"
)

// FindUserByEmail mencari user berdasarkan email
func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// CreateUser menambahkan user baru ke database
func CreateUser(user models.User) error {
	return config.DB.Create(&user).Error
}

// FindUserByID mencari user berdasarkan ID
func FindUserByID(userID uint) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, userID).Error
	return user, err
}
