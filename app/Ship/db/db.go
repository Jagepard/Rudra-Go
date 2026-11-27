package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func Init() {
	var err error
	Conn, err = sql.Open("mysql", "jagepard:password@tcp(127.0.0.1:3306)/todo_db?parseTime=true")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err := Conn.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	Conn.SetMaxOpenConns(25)
	Conn.SetMaxIdleConns(25)

	log.Println("Database connected")
}
