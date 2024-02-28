package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	if err != nil {
		fmt.Fprintf(w, "Could not render view: %v", err)
		log.Printf("Something went wrong: %v", err)
	}

	err = files.Execute(w, data)

	if err != nil {
		fmt.Fprintf(w, "Something went wrong: %v", err)
		log.Printf("Something went wrong: %v", err)
	}

}

func (v *View) RenderPartial(name string, data any) template.HTML {
	files, err := template.New(name + ".partial.html").ParseFiles(v.BasePath + "/partials/" + name + ".partial.html")

	if err != nil {
		log.Fatalf("Could not parse partials: %v", err)
		log.Printf("Something went wrong: %v", err)
	}

	writer := new(bytes.Buffer)

	err = files.Execute(writer, data)

	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
		log.Printf("Something went wrong: %v", err)
	}

	return template.HTML(writer.String())
}
