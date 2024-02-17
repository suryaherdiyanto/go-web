package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/suryaherdiyanto/go-web/src/config"
)

var appConfig config.AppConfig

func main() {
	appConfig = config.AppConfig{AppEnv: "development", UseCache: true}
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.HttpOnly = true
	session.Cookie.Path = "/"
	session.Cookie.Secure = (appConfig.AppEnv == "production")
	session.Cookie.SameSite = http.SameSiteLaxMode

	appConfig.Session = session

	router := NewRouter(&appConfig)

	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
