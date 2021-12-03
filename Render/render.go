package Render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	functions = make(template.FuncMap)
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error in getting the template cache.." + err.Error())
	}

	parsedTemplate, err := template.ParseFiles("C:/Users/91910/Documents/goworkspace/src/github.com/Webapp-New/templates/" + tmpl)

	if err != nil {
		fmt.Println("Template parsing error :" + err.Error())
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Parse error " + err.Error())
	}
}

func RenderTemplateTest(w http.ResponseWriter) (mych map[string]*template.Template, err error) {

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
