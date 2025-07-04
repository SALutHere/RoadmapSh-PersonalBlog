package handlers

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func Render(w http.ResponseWriter, r *http.Request, mode, page string, data map[string]any) {
	base := filepath.Join("static", "templates", "base.html")
	content := filepath.Join("static", "templates", mode, page+".html")

	tmpl, err := template.ParseFiles(base, content)
	if err != nil {
		http.Error(w, "template error: "+err.Error(), http.StatusInternalServerError)
	}

	if data == nil {
		data = make(map[string]any)
	}
	data["ActiveMode"] = mode

	if r.Header.Get("HX-Request") == "true" {
		tmpl.ExecuteTemplate(w, "content", data)
	} else {
		tmpl.ExecuteTemplate(w, "base", data)
	}
}
