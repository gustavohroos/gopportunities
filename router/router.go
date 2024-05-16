package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router
	r := gin.Default()

	// Initialize the routes
	initializeRoutes(r)

	// Run the router
	r.Run()
}
