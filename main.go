package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/domain/pomodoro"
	"github.com/mpawlak2/pomodoro/infrastructure/repository"
)

func main() {
	db, err := sql.Open("sqlite3", "pomodoro.db")
	if err != nil {
		panic(err)
	}
	repository.InitializeSqlLiteDB(db)
	defer db.Close()

	pomo := pomodoro.NewPomodoro(25 * time.Minute)
	pomo.Start()

	repo := repository.NewSqlLitePomodoroRepository(db)
	repo.Create(pomo)
}
