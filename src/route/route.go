package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/suryaherdiyanto/go-web/src/app"
	"github.com/suryaherdiyanto/go-web/src/middleware"
)

func NewRouter(app *app.App) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.NosurfCsrf)

	mux.Get("/", app.Home)
	mux.Post("/", app.HomePost)

	return mux
}
