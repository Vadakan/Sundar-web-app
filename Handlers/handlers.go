package Handlers

import (
	"net/http"

	"github.com/Webapp-New/Render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	Render.RenderTemplate(w, "about.page.html")
}
