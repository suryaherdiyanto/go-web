package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/suryaherdiyanto/go-web/pkg/config"
)

func main() {
	appConfig := config.AppConfig{AppEnv: "development", UseCache: true}

	router := NewRouter(&appConfig)

	fmt.Println("Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
