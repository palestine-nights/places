package main

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Application struct {
	Redis      *redis.Client
	Router     *mux.Router
	Port       string
	Experation time.Duration
}

type App = Application

func GetApp() *App {

	app := App{
		Redis:      initRedis(),
		Router:     initRouter(),
		Port:       getEnv("PORT", "8080"),
		Experation: time.Hour,
	}

	return &app
}

func (app *App) GetHandler() *http.Handler {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	handler := handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(app.Router)

	return &handler
}

func (app *App) getAddress() string {
	return "localhost:" + app.Port
}
