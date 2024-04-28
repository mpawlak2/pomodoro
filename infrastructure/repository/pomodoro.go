package repository

import (
	"database/sql"
	"time"

	"github.com/mpawlak2/pomodoro/domain/pomodoro"
)

type pomodoroDTO struct {
	ID              string
	PlannedDuration int
	Status          string
	StartTime       string
}

type SqlLitePomodoroRepository struct {
	db *sql.DB
}

func (r *SqlLitePomodoroRepository) Create(p *pomodoro.Pomodoro) error {
	dto := pomodoroDTO{
		ID:              p.ID,
		PlannedDuration: int(p.PlannedDuration),
		Status:          string(p.Status),
		StartTime:       p.StartTime.Format(time.RFC3339Nano),
	}

	_, err := r.db.Exec("INSERT INTO pomodoro (id, duration, status, start_time) VALUES (?, ?, ?, ?)", dto.ID, dto.PlannedDuration, dto.Status, dto.StartTime)
	return err
}

func (r *SqlLitePomodoroRepository) FindByID(id string) *pomodoro.Pomodoro {
	var dto pomodoroDTO

	err := r.db.QueryRow("SELECT id, duration, status, start_time FROM pomodoro WHERE id = ?", id).Scan(&dto.ID, &dto.PlannedDuration, &dto.Status, &dto.StartTime)
	if err != nil {
		return nil
	}

	startTime, err := time.Parse(time.RFC3339Nano, dto.StartTime)
	if err != nil {
		return nil
	}

	return &pomodoro.Pomodoro{
		ID:              dto.ID,
		PlannedDuration: time.Duration(dto.PlannedDuration),
		StartTime:       startTime,
		Status:          pomodoro.Status(dto.Status),
	}
}

func NewSqlLitePomodoroRepository(db *sql.DB) *SqlLitePomodoroRepository {
	return &SqlLitePomodoroRepository{
		db: db,
	}
}
