package app

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The app env: %s\n", a.Config.AppEnv)
	fmt.Printf("The use cache: %t\n", a.Config.UseCache)

	data := make(map[string]string)
	data["token"] = nosurf.Token(r)
	a.View.Render(w, "home", nil)
}

func (a *App) HomePost(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["name"] = r.FormValue("name")
	a.View.Render(w, "home", data)
}
