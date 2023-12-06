package routes

import (
	"github.com/epousa/ginFrameworkPractise/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		api.POST("/signup", handlers.SignUp)
		api.POST("/signin", handlers.SignIn)
		api.GET("/list-users", handlers.GetUsers)
	}

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/apod", handlers.GetAPIData)
	router.NoRoute(handlers.GetNoRoute)

	return router
}
