package repository

import (
	"database/sql"
	"time"

	"Rudra-Go/app/Containers/Auth/entity"
	"Rudra-Go/app/Ship/db"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.Conn}
}

// FindByLogin — поиск пользователя по логину. Возвращает sql.ErrNoRows, если не найден.
func (r *UserRepository) FindByLogin(login string) (*entity.User, error) {
	query := "SELECT id, login, password, role, created_at FROM users WHERE login = ? LIMIT 1"
	row := r.db.QueryRow(query, login)

	var u entity.User
	var createdAt time.Time
	err := row.Scan(&u.ID, &u.Login, &u.Password, &u.Role, &createdAt)
	if err != nil {
		return nil, err // sql.ErrNoRows или ошибка БД
	}
	u.CreatedAt = createdAt
	return &u, nil
}

// Create — создание пользователя. Пароль должен быть уже хэширован!
func (r *UserRepository) Create(login, passwordHash, role string) error {
	query := "INSERT INTO users (login, password, role) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, login, passwordHash, role)
	return err
}
