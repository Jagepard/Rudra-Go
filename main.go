package main

import (
	"log"
	"net/http"

	todo "Rudra-Go/app/Containers/Todo"
	ship "Rudra-Go/app/Ship"
	"Rudra-Go/app/Ship/db"
)

func main() {
	db.Init()

	app := ship.NewApp()

	containers := []ship.Container{
		todo.NewTodoContainer(),
		// admin.NewAdminContainer(),
		// auth.NewAuthContainer(),
	}

	app.Boot(containers)

	log.Println("Rudra-Go (Porto) started on :8080")
	http.ListenAndServe(":8080", app.Mux)
}
