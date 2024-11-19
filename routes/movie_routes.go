package routes

import (
	"movie-festival-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine) {
	movies := router.Group("/movies")
	{
		movies.GET("/", controllers.ListMovies)              // List movies with pagination
		movies.GET("/search", controllers.SearchMovies)      // Search movies by title/description/artists/genres
		movies.POST("/:id/view", controllers.TrackMovieView) // Track movie viewership
	}
}
