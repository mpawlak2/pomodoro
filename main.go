package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/infrastructure/repository"
)

func main() {
	db, err := sql.Open("sqlite3", "pomodoro.db")
	if err != nil {
		panic(err)
	}
	repository.InitializeSqlLiteDB(db)
	defer db.Close()
}
