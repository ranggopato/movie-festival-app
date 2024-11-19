package main

import (
	"log"
	"movie-festival-app/config"
	"movie-festival-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDatabase()

	// Setup router
	router := gin.Default()

	// Registrasi route
	routes.RegisterRoutes(router)

	// Jalankan server
	log.Println("Server berjalan di http://localhost:8080")
	router.Run(":8080")
}
