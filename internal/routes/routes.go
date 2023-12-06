package routes

import (
	"github.com/epousa/ginFrameworkPractise/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/apod", handlers.GetAPIData)
	router.NoRoute(handlers.GetNoRoute)

	return router
}
