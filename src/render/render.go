package render

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/suryaherdiyanto/go-web/src/helper"
)

type View struct {
	BasePath string
	Funcs    template.FuncMap
}

func New(path string) *View {
	return &View{
		BasePath: path,
	}
}

func (v *View) Render(w http.ResponseWriter, name string, data any) {
	files, err := template.New(name+".layout.html").Funcs(v.Funcs).ParseFiles(v.BasePath+"/"+name+".layout.html", v.BasePath+"/master.layout.html")
	helper.HandleExeption(w, err)

	err = files.Execute(w, data)
	helper.HandleExeption(w, err)
}

func (v *View) RenderPartial(name string, data any) template.HTML {
	files, err := template.New(name + ".partial.html").ParseFiles(v.BasePath + "/partials/" + name + ".partial.html")
	writer := new(bytes.Buffer)

	helper.HandleExeption(writer, err)

	err = files.Execute(writer, data)

	helper.HandleExeption(writer, err)

	return template.HTML(writer.String())
}
