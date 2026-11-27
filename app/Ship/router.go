package ship

import (
	"net/http"
	"strings"
)

// ExactHandler — для маршрутов без параметров.
// Принимает и возвращает http.Handler для единообразия с middleware.
func ExactHandler(path string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// PrefixHandler — для маршрутов с параметрами.
// Принимает и возвращает http.Handler для единообразия с middleware.
func PrefixHandler(prefix string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
