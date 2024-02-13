package main

import (
	// "um/config"
	// "um/handlers"
	"um/middleware"
	"um/routes"

	"github.com/casbin/casbin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	enforcer := casbin.NewEnforcer("models/model.conf", "models/policy.csv")

	r.Use(middleware.RBACMiddleware(enforcer))

	routes := routes.GetRoutes()
	// Iterate over routes and register them
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.Run(":8080")
}
