package ui

import (
	"fmt" // <-- ДОБАВИТЬ
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, layout, page string, data any) {
	pathLayout := "app/Containers/Todo/ui/tmpl/" + layout
	pathPage := "app/Containers/Todo/ui/tmpl/" + page

	tmpl, err := template.ParseFiles(pathLayout, pathPage)
	if err != nil {
		fmt.Println("❌ ОШИБКА ПАРСИНГА ШАБЛОНОВ:", err)
		http.Error(w, "Template Error: "+err.Error(), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		fmt.Println("❌ ОШИБКА РЕНДЕРА:", err)
		http.Error(w, "Render Error: "+err.Error(), 500)
		return
	}
}
