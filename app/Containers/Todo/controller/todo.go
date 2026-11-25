package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"Rudra-Go/app/Containers/Todo/repository"
	"Rudra-Go/app/Containers/Todo/ui"
)

type TodoController struct {
	repo *repository.TodoRepository
}

func NewTodoController() *TodoController {
	return &TodoController{
		repo: repository.NewTodoRepository(),
	}
}

func (c *TodoController) IndexAction(w http.ResponseWriter, r *http.Request) {
	todos, err := c.repo.GetAll()
	if err != nil {
		fmt.Println("❌ ОШИБКА БД:", err)
		http.Error(w, "DB Error: "+err.Error(), 500)
		return
	}

	ui.Render(w, "layout.html", "index.html", todos)
}

func (c *TodoController) CreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	task := strings.TrimSpace(r.FormValue("task"))
	if task == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c.repo.Create(task)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *TodoController) DeleteAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/todo/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	c.repo.Delete(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
