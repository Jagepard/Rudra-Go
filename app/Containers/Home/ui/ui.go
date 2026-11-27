package ui

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, layout, page string, data any) {
	tmpl, err := template.ParseFiles(
		"app/Containers/Home/ui/tmpl/"+layout,
		"app/Containers/Home/ui/tmpl/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
