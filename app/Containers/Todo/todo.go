package todo

import (
	"net/http"

	"Rudra-Go/app/Containers/Todo/controller"
	ship "Rudra-Go/app/Ship"
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
	mux.HandleFunc("/", ship.ExactHandler("/", t.ctrl.IndexAction))
	mux.HandleFunc("/todo/", ship.ExactHandler("/todo/", t.ctrl.IndexAction))
	mux.HandleFunc("/todo/create", ship.ExactHandler("/todo/create", t.ctrl.CreateAction))
	mux.HandleFunc("/todo/delete/", ship.PrefixHandler("/todo/delete/", t.ctrl.DeleteAction))
	mux.HandleFunc("/todo/hello", ship.ExactHandler("/todo/hello", t.ctrl.HelloAction))
}
