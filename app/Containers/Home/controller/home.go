package controller

import (
	"net/http"

	"Rudra-Go/app/Containers/Home/ui"
	"Rudra-Go/app/Ship"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (c *HomeController) IndexAction(w http.ResponseWriter, r *http.Request) {
	// 1. Явная проверка пути (уже сделана в ExactHandler, но на всякий случай)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 2. Проверяем авторизацию
	isAuth := ship.IsAuthorized(r)
	userID := ship.CurrentUserID(r)

	// 3. Получаем flash-сообщения
	flashErrors := ship.GetFlash(w, r, "error")
	flashSuccesses := ship.GetFlash(w, r, "success")

	// 4. Собираем данные для шаблона
	data := map[string]any{
		"IsAuth":         isAuth,
		"UserID":         userID,
		"FlashErrors":    flashErrors,
		"FlashSuccesses": flashSuccesses,
	}

	// 5. Рендерим
	ui.Render(w, "layout.html", "index.html", data)
}
