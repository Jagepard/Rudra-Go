package entity

import "time"

// User — чистая структура. Никаких методов save(), update().
type User struct {
	ID        int
	Login     string
	Password  string // Хэш пароля, никогда не plaintext
	Role      string
	CreatedAt time.Time
}
