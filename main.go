package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/application/cli"
	"github.com/mpawlak2/pomodoro/domain/pomodoro"
	"github.com/mpawlak2/pomodoro/infrastructure/repository"
)

func test() {
	db, err := sql.Open("sqlite3", "pomodoro.db")
	if err != nil {
		panic(err)
	}
	repository.InitializeSqlLiteDB(db)
	defer db.Close()

	repo := repository.NewSqlLitePomodoroRepository(db)
	pomodoroService := pomodoro.NewPomodoroService(repo)

	pomo, err := pomodoroService.CreatePomodoro(25 * time.Minute)
	if err != nil {
		panic(err)
	}
	pomo.Start()
	repo.Create(pomo)
}

func main() {
	cli.RunApplication()
}
