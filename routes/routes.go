package routes

import (
	"um/db"
	"um/handlers"
	"um/repositories"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

var userRepository repositories.UserRepository

func init() {
	db.InitDB()
	userRepository = repositories.NewGormUserRepository(db.GetDB())
}

var userRoutes = map[string]Route{
	"/signup": {
		Method:      "POST",
		Path:        "/signup",
		HandlerFunc: handlers.NewUserHandler(userRepository).Signup,
	},
	"/login": {
		Method:      "POST",
		Path:        "/login",
		HandlerFunc: handlers.NewUserHandler(userRepository).Login,
	},
	"/me": {
		Method:      "GET",
		Path:        "/me",
		HandlerFunc: handlers.NewUserHandler(userRepository).GetMe,
	},
}

func GetRoutes() map[string]Route {
	return userRoutes
}
