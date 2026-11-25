package ship

import "net/http"

type App struct {
	Mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		Mux: http.NewServeMux(),
	}
}

func (a *App) Boot(containers []Container) {
	for _, c := range containers {
		c.RegisterRoutes(a.Mux)
	}
}
