package server

import (
	"github.com/epousa/ginFrameworkPractise/internal/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/apod", handlers.GetAPIData)

	router.Run("localhost:8080")
}
