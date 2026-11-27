package home

import (
	"net/http"

	"Rudra-Go/app/Containers/Home/controller"
	"Rudra-Go/app/Ship"
)

type HomeContainer struct {
	ctrl *controller.HomeController
}

func NewHomeContainer() *HomeContainer {
	return &HomeContainer{
		ctrl: controller.NewHomeController(),
	}
}

func (h *HomeContainer) Name() string {
	return "Home"
}

func (h *HomeContainer) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", ship.ExactHandler("/",
		http.HandlerFunc(h.ctrl.IndexAction)))
}
