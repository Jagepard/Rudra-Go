package todo

import (
	"net/http"

	"Rudra-Go/app/Containers/Todo/controller"
)

type Config struct {
	MaxTasks int
}

var Cfg = Config{
	MaxTasks: 100,
}

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
	mux.HandleFunc("/", t.ctrl.IndexAction)
	mux.HandleFunc("/todo/create", t.ctrl.CreateAction)
	mux.HandleFunc("/todo/delete/", t.ctrl.DeleteAction)
}
