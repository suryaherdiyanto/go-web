package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suryaherdiyanto/go-web/src/app"
	"github.com/suryaherdiyanto/go-web/src/config"
	"github.com/suryaherdiyanto/go-web/src/helper"
	"github.com/suryaherdiyanto/go-web/src/route"
	"github.com/suryaherdiyanto/go-web/src/session"
)

var appConfig config.AppConfig

func main() {
	helper.LoadEnv("./.env")
	appConfig = config.AppConfig{AppEnv: os.Getenv("APP_ENV"), UseCache: true}
	log.Print(appConfig)
	app := app.New(&appConfig)
	app.Session = session.New(&appConfig)

	router := route.NewRouter(app)

	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
