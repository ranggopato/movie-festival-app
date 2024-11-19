package config

import (
	"log"
	"movie-festival-app/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// ConnectDatabase membuat koneksi ke database dan melakukan migrasi
func ConnectDatabase() {
	// Gunakan github.com/glebarez/sqlite sebagai driver
	database, err := gorm.Open(sqlite.Open("movie_festival.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Migrasi model ke database (membuat tabel jika belum ada)
	err = database.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Vote{},
		&models.Viewership{},
	)
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}

	DB = database
	log.Println("Berhasil terhubung ke database SQLite menggunakan github.com/glebarez/sqlite!")
}
