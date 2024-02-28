package app

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/suryaherdiyanto/go-web/src/render"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The app env: %s\n", a.Config.AppEnv)
	fmt.Printf("The use cache: %t\n", a.Config.UseCache)

	data := make(map[string]string)
	data["token"] = nosurf.Token(r)
	render.RenderView(w, "home", data)
}

func (a *App) HomePost(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["name"] = r.FormValue("name")
	render.RenderView(w, "home", data)
}
