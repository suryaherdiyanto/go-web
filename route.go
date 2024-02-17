package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/suryaherdiyanto/go-web/pkg/config"
	"github.com/suryaherdiyanto/go-web/pkg/handlers"
	"github.com/suryaherdiyanto/go-web/pkg/repository"
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
