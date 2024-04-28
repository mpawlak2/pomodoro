package repository

import (
	"database/sql"
	"time"

	"github.com/mpawlak2/pomodoro/domain/pomodoro"
)

type pomodoroDTO struct {
	ID       string
	Duration int
	Status   string
}

type SqlLitePomodoroRepository struct {
	db *sql.DB
}

func (r *SqlLitePomodoroRepository) Create(p *pomodoro.Pomodoro) error {
	dto := pomodoroDTO{
		ID:       p.ID,
		Duration: int(p.Duration),
		Status:   string(p.Status),
	}

	_, err := r.db.Exec("INSERT INTO pomodoro (id, duration, status) VALUES (?, ?, ?)", dto.ID, dto.Duration, dto.Status)
	return err
}

func (r *SqlLitePomodoroRepository) FindByID(id string) *pomodoro.Pomodoro {
	var dto pomodoroDTO

	err := r.db.QueryRow("SELECT id, duration, status FROM pomodoro WHERE id = ?", id).Scan(&dto.ID, &dto.Duration, &dto.Status)
	if err != nil {
		return nil
	}

	return &pomodoro.Pomodoro{
		ID:       dto.ID,
		Duration: time.Duration(dto.Duration),
		Status:   pomodoro.Status(dto.Status),
	}
}

func NewSqlLitePomodoroRepository(db *sql.DB) *SqlLitePomodoroRepository {
	return &SqlLitePomodoroRepository{
		db: db,
	}
}
