package test

import (
	"github.com/go-chi/chi/v5"
	"github.com/suryaherdiyanto/go-web/src/app"
	"github.com/suryaherdiyanto/go-web/src/config"
	"github.com/suryaherdiyanto/go-web/src/route"
	"github.com/suryaherdiyanto/go-web/src/session"
)

var singleton *app.App
var routes *chi.Mux
var appConfig *config.AppConfig

func Setup() {
	appConfig = &config.AppConfig{AppEnv: "testing", UseCache: true}
	singleton = app.New(appConfig)
	singleton.Session = session.New(appConfig)
	routes = route.NewRouter(singleton)
}
