package repository

import (
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/domain/pomodoro"
)

func TestSqlLitePomodoroRepository(t *testing.T) {
	defer os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf("Error opening SQLite database: %v", err)
	}
	defer db.Close()

	repo := NewSqlLitePomodoroRepository(db)

	if repo.db == nil {
		t.Errorf("Expected db to be set, but got %v", repo.db)
	}
}

func TestPersistPomodoro(t *testing.T) {
	defer os.Remove("./test.db") // todo: make it possible to run tests in parallel A: it could be done by using a different database file for each test or by using a different table for each test

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf("Error opening SQLite database: %v", err)
	}
	InitializeSqlLiteDB(db)
	defer db.Close()

	repo := NewSqlLitePomodoroRepository(db)
	pomodoro := pomodoro.NewPomodoro(25 * time.Minute)
	pomodoro.Start()

	err = repo.Create(pomodoro)
	if err != nil {
		t.Errorf("Error persisting pomodoro: %v", err)
	}

	pomo := repo.FindByID(pomodoro.ID)
	if pomo == nil {
		t.Errorf("Expected to find pomodoro, but got nil")
	}

	if pomo.PlannedDuration != pomodoro.PlannedDuration {
		t.Errorf("Expected pomodoro duration to be %v, but got %v", pomodoro.PlannedDuration, pomo.PlannedDuration)
	}

	if !pomo.StartTime.Equal(pomodoro.StartTime) {
		t.Errorf("Expected pomodoro start time to be %v, but got %v", pomodoro.StartTime, pomo.StartTime)
	}

	if pomo.ID != pomodoro.ID {
		t.Errorf("Expected pomodoro ID to be %v, but got %v", pomodoro.ID, pomo.ID)
	}

	if pomo.Status != pomodoro.Status {
		t.Errorf("Expected pomodoro status to be %v, but got %v", pomodoro.Status, pomo.Status)
	}
}
