package ship

import (
	"net/http"
	"strings"
)

// ExactHandler — для маршрутов без параметров.
// Проверяет точное совпадение пути. Если путь не совпадает — возвращает 404.
func ExactHandler(path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}
}

// PrefixHandler — для маршрутов с параметрами (например, /todo/delete/123).
// Проверяет, что путь начинается с префикса.
// Handler сам извлекает параметр из r.URL.Path.
func PrefixHandler(prefix string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	}
}
