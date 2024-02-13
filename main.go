package main

import (
	// "um/config"
	// "um/handlers"
	// "um/middleware"
	"um/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Add middleware (if applicable)
	// r.Use(middleware.Cors())
	// r.Use(middleware.JWTAuth())

	routes := routes.GetRoutes()
	// Iterate over routes and register them
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.Run(":8080")
}
