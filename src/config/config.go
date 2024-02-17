package config

import "github.com/alexedwards/scs/v2"

type AppConfig struct {
	AppEnv   string
	UseCache bool
	Session  *scs.SessionManager
}
