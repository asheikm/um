package routes

import (
	"um/db"
	"um/handlers"
	"um/services"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

var userService services.UserService

func init() {
	db.InitDB()
	userService = services.NewGormUserService(db.GetDB())
}

var userRoutes = map[string]Route{
	"/signup": {
		Method:      "POST",
		Path:        "/signup",
		HandlerFunc: handlers.NewUserHandler(userService).Signup,
	},
	"/login": {
		Method:      "POST",
		Path:        "/login",
		HandlerFunc: handlers.NewUserHandler(userService).Login,
	},
	"/me": {
		Method:      "GET",
		Path:        "/me",
		HandlerFunc: handlers.NewUserHandler(userService).GetMe,
	},
}

func GetRoutes() map[string]Route {
	return userRoutes
}
