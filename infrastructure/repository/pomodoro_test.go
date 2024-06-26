package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/domain/pomodoro"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("sqlite3", "./test.db") // todo: is there a way to create a global database for all integration tests?
	if err != nil {
		log.Fatal(err)
	}
	InitializeSqlLiteDB(db)
	defer db.Close()

	code := m.Run()
	os.Exit(code)
}

func TestSqlLitePomodoroRepository(t *testing.T) {
	repo := NewSqlLitePomodoroRepository(db)

	if repo.db == nil {
		t.Errorf("Expected db to be set, but got %v", repo.db)
	}
}

func TestPersistPomodoro(t *testing.T) {
	repo := NewSqlLitePomodoroRepository(db)
	pomodoro := pomodoro.NewPomodoro(25 * time.Minute)
	pomodoro.Start()

	err := repo.Create(pomodoro)
	if err != nil {
		t.Errorf("Error persisting pomodoro: %v", err)
	}

	pomo, err := repo.FindByID(pomodoro.ID)
	if err != nil {
		t.Errorf("Error finding pomodoro: %v", err)
	}

	if pomo == nil {
		t.Errorf("Expected to find pomodoro, but got nil")
		panic("unreachable") // dirty fix for staticcheck SA5011
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

func TestFindActivePomodoro(t *testing.T) {
	db.Exec("DELETE FROM pomodoro") // todo: how to handle this? How to ensure the tests can be executed in parallel?
	repo := NewSqlLitePomodoroRepository(db)
	pomodoro := pomodoro.NewPomodoro(25 * time.Minute)
	pomodoro.Start()

	err := repo.Create(pomodoro)
	if err != nil {
		t.Errorf("Error persisting pomodoro: %v", err)
	}

	activePomodoro, err := repo.FindActive()
	if err != nil {
		t.Errorf("Error finding active pomodoro: %v", err)
	}

	if activePomodoro == nil {
		t.Errorf("Expected to find active pomodoro, but got nil")
		panic("unreachable") // dirty fix for staticcheck SA5011
	}

	if activePomodoro.ID != pomodoro.ID {
		t.Errorf("Expected active pomodoro ID to be %v, but got %v", pomodoro.ID, activePomodoro.ID)
	}
}

func TestFindAllPomodoros(t *testing.T) {
	repo := NewSqlLitePomodoroRepository(db)
	pomodoro := pomodoro.NewPomodoro(25 * time.Minute)
	pomodoro.Start()

	err := repo.Create(pomodoro)
	if err != nil {
		t.Errorf("Error persisting pomodoro: %v", err)
	}

	pomodoros, err := repo.FindAll()
	if err != nil {
		t.Errorf("Error finding all pomodoros: %v", err)
	}

	if len(pomodoros) < 1 {
		t.Errorf("Expected to find at least 1 pomodoro, but got %v", len(pomodoros))
	}
}

func TestUpdatePomodoro(t *testing.T) {
	repo := NewSqlLitePomodoroRepository(db)
	pomodoro := pomodoro.NewPomodoro(25 * time.Minute)
	pomodoro.Start()

	err := repo.Create(pomodoro)
	if err != nil {
		t.Errorf("Error persisting pomodoro: %v", err)
	}

	pomodoro.Finish("Test note")
	expectedFinishedTime := pomodoro.FinishTime
	err = repo.Update(pomodoro)
	if err != nil {
		t.Errorf("Error updating pomodoro: %v", err)
	}

	pomo, err := repo.FindByID(pomodoro.ID)
	if err != nil {
		t.Errorf("Error finding pomodoro: %v", err)
	}

	if pomo == nil {
		t.Errorf("Expected to find pomodoro, but got nil")
		panic("unreachable") // dirty fix for staticcheck SA5011
	}

	if pomo.Status != pomodoro.Status {
		t.Errorf("Expected pomodoro status to be %v, but got %v", pomodoro.Status, pomo.Status)
	}

	if pomo.Note != pomodoro.Note {
		t.Errorf("Expected pomodoro note to be %v, but got %v", pomodoro.Note, pomo.Note)
	}

	if pomo.FinishTime.IsZero() {
		t.Errorf("Expected finish time to be set, but got zero")
	}

	if !pomo.FinishTime.Equal(expectedFinishedTime) {
		t.Errorf("Expected finish time to be %v, but got %v", expectedFinishedTime, pomo.FinishTime)
	}
}
