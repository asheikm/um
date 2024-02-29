package main

import (
	// "flag"
	"fmt"
	"os"
	"time"
	"um/handlers"
	"um/middleware"
	"um/routes"

	"github.com/casbin/casbin"

	"um/db"
	"um/services"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if raw != "" {
			path = path + "?" + raw
		}

		logger := log.With().
			Str("method", c.Request.Method).
			Str("path", path).
			Str("ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str("status", fmt.Sprintf("%d", c.Writer.Status())).
			Dur("latency", latency).
			Logger()

		if len(c.Errors) > 0 {
			// Append error to log entry
			logger.Error().Msg(c.Errors.String())
		} else {
			logger.Info()
		}
	}
}

func main() {
	log.Info().Msg("UM gin router initialized")
	r := gin.New()
	log.Info().Msg("UM db initialized")
	db.InitDB()

	enforcer := casbin.NewEnforcer("models/model.conf", "models/policy.csv")

	r.Use(middleware.RBACMiddleware(enforcer))
	r.Use(ginLogger())

	db := db.GetDB()

	userService := services.NewGormUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	v1 := r.Group("/api/v1")

	routes := routes.GetRoutes(userHandler)
	// Iterate over routes and register them
	for _, route := range routes {
		v1.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.Run(":8080")
}
