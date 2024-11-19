package controllers

import (
	"movie-festival-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListMovies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	movies, err := services.ListMovies(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan daftar film"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func SearchMovies(c *gin.Context) {
	title := c.Query("title")
	description := c.Query("description")
	artists := c.Query("artists")
	genres := c.Query("genres")

	movies, err := services.SearchMovies(title, description, artists, genres)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencari film"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func TrackMovieView(c *gin.Context) {
	movieID, _ := strconv.Atoi(c.Param("id"))
	err := services.TrackMovieView(uint(movieID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melacak viewership"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Viewership berhasil dilacak"})
}

func TrackViewershipHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID film tidak valid"})
		return
	}

	var body struct {
		Duration int `json:"duration" binding:"required"` // Durasi dalam detik
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	err = services.TrackMovieViewership(userID.(uint), uint(movieID), body.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tontonan berhasil dicatat"})
}
