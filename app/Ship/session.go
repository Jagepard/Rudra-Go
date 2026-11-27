package ship

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Store — глобальный store для сессий. В продакшене используйте CookieStore с секретом из env.
var Store = sessions.NewCookieStore([]byte("your-secret-key-change-in-production"))

// SetFlash — устанавливает flash-сообщение в сессию
func SetFlash(w http.ResponseWriter, r *http.Request, name, value string) {
	session, _ := Store.Get(r, "flash-session")
	session.AddFlash(value, name)
	session.Save(r, w)
}

// GetFlash — получает flash-сообщение и удаляет его из сессии
func GetFlash(w http.ResponseWriter, r *http.Request, name string) []interface{} {
	session, _ := Store.Get(r, "flash-session")
	flashes := session.Flashes(name)
	session.Save(r, w) // Теперь w доступен и корректен
	return flashes
}

// RedirectPRG — Post/Redirect/Get паттерн
func RedirectPRG(w http.ResponseWriter, r *http.Request, to string) {
	http.Redirect(w, r, to, http.StatusSeeOther)
}

// CurrentUserID возвращает ID пользователя, если он авторизован.
// Если нет — возвращает 0.
func CurrentUserID(r *http.Request) int {
	session, _ := Store.Get(r, "auth-session")
	// Типизация важна: gorilla/sessions хранит всё как interface{}
	if id, ok := session.Values["user_id"].(int); ok {
		return id
	}
	return 0
}

// IsAuthorized — удобная обертка для проверки факта авторизации.
func IsAuthorized(r *http.Request) bool {
	return CurrentUserID(r) > 0
}
