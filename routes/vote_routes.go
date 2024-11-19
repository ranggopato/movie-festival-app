package routes

import (
	"movie-festival-app/controllers"
	"movie-festival-app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterVoteRoutes(router *gin.Engine) {
	votes := router.Group("/votes")
	votes.Use(middlewares.AuthMiddleware()) // Hanya pengguna yang sudah login yang dapat mem-vote
	{
		votes.POST("/:movie_id", controllers.VoteMovieHandler)     // Vote a movie
		votes.DELETE("/:movie_id", controllers.UnvoteMovieHandler) // Unvote a movie
		votes.GET("/", controllers.ListUserVotedMoviesHandler)     // List all voted movies by user
	}
}
