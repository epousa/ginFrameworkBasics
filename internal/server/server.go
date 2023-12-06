package server

import "github.com/epousa/ginFrameworkPractise/internal/routes"

func StartServer() {
	router := routes.SetRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
