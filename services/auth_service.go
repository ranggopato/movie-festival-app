package services

import (
	"errors"
	"movie-festival-app/models"
	"movie-festival-app/repositories"
	"movie-festival-app/utils"
)

// RegisterUser mendaftarkan user baru
func RegisterUser(user models.User) error {
	// Periksa apakah email sudah digunakan
	existingUser, _ := repositories.FindUserByEmail(user.Email)
	if existingUser.ID != 0 {
		return errors.New("email sudah terdaftar")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return repositories.CreateUser(user)
}

// LoginUser memvalidasi kredensial login
func LoginUser(email, password string) (string, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil || user.ID == 0 {
		return "", errors.New("email atau password salah")
	}

	// Validasi password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("email atau password salah")
	}

	// Buat token JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID mengambil user berdasarkan ID
func GetUserByID(userID uint) (models.User, error) {
	return repositories.FindUserByID(userID)
}
