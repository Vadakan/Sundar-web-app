package Render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	functions = make(template.FuncMap)
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache(w)
	if err != nil {
		log.Fatal(err.Error())
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err.Error())
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateTemplateCache(w http.ResponseWriter) (mych map[string]*template.Template, err error) {

	mycache := make(map[string]*template.Template)

	pages, err := filepath.Glob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.page.html")

	if err != nil {
		return mycache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return mycache, err
		}
		matches, err := filepath.Glob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.layout.html")
		if err != nil {
			return mycache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/*.layout.html")
			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts
	}
	return mycache, nil
}
