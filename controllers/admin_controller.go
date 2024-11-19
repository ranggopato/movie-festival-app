package controllers

import (
	"movie-festival-app/models"
	"movie-festival-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan film"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Film berhasil ditambahkan"})
}

func UpdateMovie(c *gin.Context) {
	movieID, _ := strconv.Atoi(c.Param("id"))
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateMovie(uint(movieID), movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui film"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Film berhasil diperbarui"})
}

func GetMostViewedMovie(c *gin.Context) {
	movie, err := services.GetMostViewedMovie()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan film dengan view terbanyak"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func GetMostViewedGenres(c *gin.Context) {
	genres, err := services.GetMostViewedGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan genre dengan view terbanyak"})
		return
	}

	c.JSON(http.StatusOK, genres)
}
