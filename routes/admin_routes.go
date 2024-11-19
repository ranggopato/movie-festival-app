package routes

import (
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	admin.Use(middlewares.AuthMiddleware()) // Memastikan hanya admin yang bisa mengakses
	{
		admin.POST("/movies", controllers.CreateMovie)                         // Create movie
		admin.PUT("/movies/:id", controllers.UpdateMovie)                      // Update movie
		admin.GET("/movies/most-viewed", controllers.GetMostViewedMovie)       // Most viewed movie
		admin.GET("/genres/most-viewed", controllers.GetMostViewedGenres)      // Most viewed genre
		admin.GET("/movies/most-voted", controllers.GetMostVotedMoviesHandler) // Most voted movie
		admin.POST("/:movie_id/view", controllers.TrackViewershipHandler)
	}
}
