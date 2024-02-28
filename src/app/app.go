package app

import (
	"database/sql"

	"github.com/alexedwards/scs/v2"
	"github.com/suryaherdiyanto/go-web/src/config"
)

type App struct {
	Config  *config.AppConfig
	Session *scs.SessionManager
	DB      *sql.DB
}

func New(config *config.AppConfig) *App {
	return &App{
		Config: config,
	}
}
