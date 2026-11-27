package controller

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"Rudra-Go/app/Containers/Auth/repository"
	"Rudra-Go/app/Containers/Auth/ui"
	ship "Rudra-Go/app/Ship"
)

type AuthController struct {
	repo *repository.UserRepository
}

func NewAuthController() *AuthController {
	return &AuthController{
		repo: repository.NewUserRepository(),
	}
}

// LoginAction — отображает форму входа (GET) или обрабатывает вход (POST)
func (c *AuthController) LoginAction(w http.ResponseWriter, r *http.Request) {
	log.Printf("AUTH: LoginAction called, method=%s, path=%s", r.Method, r.URL.Path)

	if r.Method == "POST" {
		c.handleLogin(w, r)
		return
	}

	ui.Render(w, "layout.html", "login.html", nil)
}

func (c *AuthController) handleLogin(w http.ResponseWriter, r *http.Request) {
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))

	if login == "" || password == "" {
		log.Printf("AUTH: empty credentials, redirecting")
		ship.SetFlash(w, r, "error", "Логин и пароль обязательны")
		ship.RedirectPRG(w, r, "/auth/login")
		return
	}

	// Ищем пользователя
	user, err := c.repo.FindByLogin(login)
	if err != nil {
		ship.SetFlash(w, r, "error", "Неверный логин или пароль")
		ship.RedirectPRG(w, r, "/")
		return
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("AUTH: bcrypt MISMATCH, err=%v", err)
		ship.SetFlash(w, r, "error", "Неверный логин или пароль")
		ship.RedirectPRG(w, r, "/")
		return
	}

	log.Printf("AUTH: login SUCCESS for %s", login)

	// Создаём сессию
	session, _ := ship.Store.Get(r, "auth-session")
	session.Values["user_id"] = user.ID
	session.Values["login"] = user.Login
	session.Save(r, w)

	ship.SetFlash(w, r, "success", "Добро пожаловать, "+user.Login)
	ship.RedirectPRG(w, r, "/")
}

// LogoutAction — выход из системы
func (c *AuthController) LogoutAction(w http.ResponseWriter, r *http.Request) {
	session, _ := ship.Store.Get(r, "auth-session")
	session.Values["user_id"] = nil
	session.Save(r, w)

	ship.RedirectPRG(w, r, "/auth/login")
}

// RegisterAction — регистрация нового пользователя
func (c *AuthController) RegisterAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.handleRegister(w, r)
		return
	}

	ui.Render(w, "layout.html", "register.html", nil)
}

func (c *AuthController) handleRegister(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password_confirm")

	// Валидация
	if login == "" || password == "" {
		ship.SetFlash(w, r, "error", "Логин и пароль обязательны")
		ship.RedirectPRG(w, r, "/auth/register")
		return
	}

	if password != passwordConfirm {
		ship.SetFlash(w, r, "error", "Пароли не совпадают")
		ship.RedirectPRG(w, r, "/auth/register")
		return
	}

	if len(password) < 6 {
		ship.SetFlash(w, r, "error", "Пароль должен быть не менее 6 символов")
		ship.RedirectPRG(w, r, "/auth/register")
		return
	}

	// Хэшируем пароль (bcrypt, cost=10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ship.SetFlash(w, r, "error", "Ошибка сервера")
		ship.RedirectPRG(w, r, "/auth/register")
		return
	}

	// Создаём пользователя
	if err := c.repo.Create(login, string(hashedPassword), "user"); err != nil {
		ship.SetFlash(w, r, "error", "Пользователь уже существует")
		ship.RedirectPRG(w, r, "/auth/register")
		return
	}

	ship.SetFlash(w, r, "success", "Регистрация успешна. Войдите в систему.")
	ship.RedirectPRG(w, r, "/auth/login")
}
