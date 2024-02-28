package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/suryaherdiyanto/go-web/src/config"
)

func New(config *config.AppConfig) *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.HttpOnly = true
	session.Cookie.Path = "/"
	session.Cookie.Secure = (config.AppEnv == "production")
	session.Cookie.SameSite = http.SameSiteLaxMode

	return session
}
