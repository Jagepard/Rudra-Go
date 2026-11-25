package repository

import (
	"Rudra-Go/app/Containers/Todo/entity"
	"Rudra-Go/app/Ship/db" // Глобальный пул из Ship
	"database/sql"
)

type TodoRepository struct {
	db *sql.DB // В реальном проекте лучше передавать через конструктор, но для простоты берем из Ship
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{db: db.Conn}
}

func (r *TodoRepository) GetAll() ([]entity.Todo, error) {
	rows, err := r.db.Query("SELECT id, task, completed FROM todos ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []entity.Todo
	for rows.Next() {
		var t entity.Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, rows.Err()
}

func (r *TodoRepository) Create(task string) error {
	_, err := r.db.Exec("INSERT INTO todos (task) VALUES (?)", task)
	return err
}

func (r *TodoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
