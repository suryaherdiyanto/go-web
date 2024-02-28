package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/suryaherdiyanto/go-web/src/app"
	"github.com/suryaherdiyanto/go-web/src/handlers"
	"github.com/suryaherdiyanto/go-web/src/middleware"
	"github.com/suryaherdiyanto/go-web/src/repository"
)

func NewRouter(a *app.App) *chi.Mux {
	appRepository := repository.New(a.Config)
	baseHandler := handlers.NewBaseHandler(appRepository)

	mux := chi.NewRouter()
	mux.Use(middleware.NosurfCsrf)

	mux.Get("/", baseHandler.Home)
	mux.Post("/", baseHandler.HomePost)

	return mux
}
