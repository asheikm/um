package routes

import (
	"um/db"
	"um/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func init() {
	db.InitDB()
}

func GetRoutes(userHandler *handlers.UserHandler) []Route {
	userRoutes := []Route{
		{
			Method:      "POST",
			Path:        "/signup",
			HandlerFunc: userHandler.Signup,
		},
		{
			Method:      "POST",
			Path:        "/login",
			HandlerFunc: userHandler.Login,
		},
		{
			Method:      "GET",
			Path:        "/me",
			HandlerFunc: userHandler.GetMe,
		},
	}
	return userRoutes
}
