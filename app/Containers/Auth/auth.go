package auth

import (
	"net/http"

	"Rudra-Go/app/Containers/Auth/controller"
	"Rudra-Go/app/Containers/Auth/middleware"
	"Rudra-Go/app/Ship"
)

type Config struct {
	SessionMaxAge int // секунды
}

var Cfg = Config{
	SessionMaxAge: 86400, // 24 часа
}

type AuthContainer struct {
	ctrl *controller.AuthController
}

func NewAuthContainer() *AuthContainer {
	return &AuthContainer{
		ctrl: controller.NewAuthController(),
	}
}

func (a *AuthContainer) Name() string {
	return "Auth"
}

func (a *AuthContainer) RegisterRoutes(mux *http.ServeMux) {
	// Публичные маршруты
	mux.Handle("/auth/login", ship.ExactHandler("/auth/login",
		http.HandlerFunc(a.ctrl.LoginAction)))
	mux.Handle("/auth/register", ship.ExactHandler("/auth/register",
		http.HandlerFunc(a.ctrl.RegisterAction)))

	// Защищённый маршрут
	mux.Handle("/auth/logout", ship.ExactHandler("/auth/logout",
		middleware.AuthRequired(http.HandlerFunc(a.ctrl.LogoutAction))))
}
