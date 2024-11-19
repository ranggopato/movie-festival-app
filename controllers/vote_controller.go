package controllers

import (
	"movie-festival-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VoteMovieHandler menangani permintaan untuk memberikan suara pada sebuah film
func VoteMovieHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	movieID, err := strconv.ParseUint(c.Param("movie_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID film tidak valid"})
		return
	}

	err = services.VoteMovie(uint(userID.(uint)), uint(movieID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote berhasil ditambahkan"})
}

// UnvoteMovieHandler menangani permintaan untuk mencabut suara pada sebuah film
func UnvoteMovieHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	movieID, err := strconv.ParseUint(c.Param("movie_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID film tidak valid"})
		return
	}

	err = services.UnvoteMovie(uint(userID.(uint)), uint(movieID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote berhasil dicabut"})
}

// ListUserVotedMoviesHandler menangani permintaan untuk melihat semua film yang di-vote oleh user
func ListUserVotedMoviesHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	movies, err := services.ListUserVotedMovies(uint(userID.(uint)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar film"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// GetMostVotedMoviesHandler menangani permintaan untuk mendapatkan film paling banyak di-vote
func GetMostVotedMoviesHandler(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter limit tidak valid"})
		return
	}

	movies, err := services.GetMostVotedMovies(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar film"})
		return
	}

	c.JSON(http.StatusOK, movies)
}
