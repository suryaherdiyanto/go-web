package app

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/suryaherdiyanto/go-web/src/config"
	"github.com/suryaherdiyanto/go-web/src/render"
)

type App struct {
	Config  *config.AppConfig
	Session *scs.SessionManager
	DB      *sql.DB
	View    *render.View
}

func New(config *config.AppConfig) *App {
	return &App{
		Config: config,
		View:   render.New(config.BaseViewPath),
	}
}
