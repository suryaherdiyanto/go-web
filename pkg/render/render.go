package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var funcs = template.FuncMap{
	"RenderPartial": RenderPartial,
}

func RenderView(w http.ResponseWriter, name string, data any) {
	files, err := template.New(name+".layout.html").Funcs(funcs).ParseFiles("./views/"+name+".layout.html", "./views/master.layout.html")

	if err != nil {
		fmt.Fprintf(w, "Could not render view: %v", err)
	}

	err = files.Execute(w, data)

	if err != nil {
		fmt.Fprintf(w, "Something went wrong: %v", err)
	}

}

func RenderPartial(name string, data any) template.HTML {
	files, err := template.ParseFiles("./views/partials/" + name + ".partial.html")

	if err != nil {
		log.Fatalf("Could not parse partials: %v", err)
	}

	writer := new(bytes.Buffer)

	err = files.Execute(writer, data)

	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	return template.HTML(writer.String())
}
