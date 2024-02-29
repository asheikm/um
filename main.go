package main

import (
	"um/handlers"
	"um/middleware"
	"um/routes"

	"github.com/casbin/casbin"

	"um/db"
	"um/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	enforcer := casbin.NewEnforcer("models/model.conf", "models/policy.csv")

	r.Use(middleware.RBACMiddleware(enforcer))

	db := db.GetDB()

	userService := services.NewGormUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	routes := routes.GetRoutes(userHandler)
	// Iterate over routes and register them
	for _, route := range routes {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.Run(":8080")
}
