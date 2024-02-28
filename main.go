package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryaherdiyanto/go-web/src/app"
	"github.com/suryaherdiyanto/go-web/src/config"
	"github.com/suryaherdiyanto/go-web/src/session"
)

var appConfig config.AppConfig

func main() {
	appConfig = config.AppConfig{AppEnv: "development", UseCache: true}
	app := app.New(&appConfig)
	app.Session = session.New(&appConfig)

	router := NewRouter(&appConfig)

	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
