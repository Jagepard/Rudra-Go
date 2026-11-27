package todo

import (
	"net/http"

	"Rudra-Go/app/Containers/Auth/middleware"
	"Rudra-Go/app/Containers/Todo/controller"
	"Rudra-Go/app/Ship"
)

type TodoContainer struct {
	ctrl *controller.TodoController
}

func NewTodoContainer() *TodoContainer {
	return &TodoContainer{
		ctrl: controller.NewTodoController(),
	}
}

func (t *TodoContainer) Name() string {
	return "Todo"
}

func (t *TodoContainer) RegisterRoutes(mux *http.ServeMux) {
	protectedMux := http.NewServeMux()

	protectedMux.Handle("/todo/", ship.ExactHandler("/todo/",
		http.HandlerFunc(t.ctrl.IndexAction)))
	protectedMux.Handle("/todo/create", ship.ExactHandler("/todo/create",
		http.HandlerFunc(t.ctrl.CreateAction)))
	protectedMux.Handle("/todo/delete/", ship.PrefixHandler("/todo/delete/",
		http.HandlerFunc(t.ctrl.DeleteAction)))
	protectedMux.Handle("/todo/hello", ship.ExactHandler("/todo/hello",
		http.HandlerFunc(t.ctrl.HelloAction)))

	mux.Handle("/todo/", middleware.AuthRequired(protectedMux))
}
