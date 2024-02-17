package handlers

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/suryaherdiyanto/go-web/pkg/render"
	"github.com/suryaherdiyanto/go-web/pkg/repository"
)

type BaseHandler struct {
	Repository *repository.AppRepository
}

func NewBaseHandler(repo *repository.AppRepository) *BaseHandler {
	return &BaseHandler{
		Repository: repo,
	}
}

func (h *BaseHandler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("The app env: %s\n", h.Repository.Config.AppEnv)
	fmt.Printf("The use cache: %t\n", h.Repository.Config.UseCache)

	data := make(map[string]string)
	data["token"] = nosurf.Token(r)
	render.RenderView(w, "home", data)
}

func (h *BaseHandler) HomePost(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["name"] = r.FormValue("name")
	render.RenderView(w, "home", data)
}
