package main

import (
	"log"
	"net/http"

	auth "Rudra-Go/app/Containers/Auth"
	home "Rudra-Go/app/Containers/Home"
	todo "Rudra-Go/app/Containers/Todo"
	ship "Rudra-Go/app/Ship"
	"Rudra-Go/app/Ship/db"
)

func main() {
	db.Init()

	db.Conn.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		login VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL DEFAULT 'user',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	app := ship.NewApp()

	containers := []ship.Container{
		todo.NewTodoContainer(),
		auth.NewAuthContainer(),
		home.NewHomeContainer(),
	}

	app.Boot(containers)

	log.Println("Rudra-Go (Porto) started on :8080")
	http.ListenAndServe(":8080", app.Mux)
}
