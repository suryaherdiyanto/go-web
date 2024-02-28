package config

type AppConfig struct {
	AppEnv       string
	AppURL       string
	UseCache     bool
	DBDriver     string
	DBURL        string
	BaseViewPath string
}
