package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/suryaherdiyanto/go-web/src/config"
	"github.com/suryaherdiyanto/go-web/src/handlers"
	"github.com/suryaherdiyanto/go-web/src/repository"
)

func NewRouter(appConfig *config.AppConfig) *chi.Mux {
	appRepository := repository.New(appConfig)
	baseHandler := handlers.NewBaseHandler(appRepository)

	mux := chi.NewRouter()
	mux.Use(NosurfCsrf)

	mux.Get("/", baseHandler.Home)
	mux.Post("/", baseHandler.HomePost)

	return mux
}
