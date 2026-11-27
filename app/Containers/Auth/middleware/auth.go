package middleware

import (
	"Rudra-Go/app/Ship"
	"net/http"
)

// AuthRequired принимает http.Handler (интерфейс), а не http.HandlerFunc.
// Это позволяет оборачивать и функции, и ServeMux, и любые другие handler'ы.
func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := ship.Store.Get(r, "auth-session")

		if session.Values["user_id"] == nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
