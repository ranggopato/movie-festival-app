package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Registrasi semua rute
	RegisterAdminRoutes(router)
	RegisterMovieRoutes(router)
	RegisterVoteRoutes(router)
	RegisterAuthRoutes(router)
}
