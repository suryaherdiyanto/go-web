package repository

import "github.com/suryaherdiyanto/go-web/src/config"

type AppRepository struct {
	Config *config.AppConfig
}

func New(appConfig *config.AppConfig) *AppRepository {
	return &AppRepository{
		Config: appConfig,
	}
}
