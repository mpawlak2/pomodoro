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
	FinishTime      sql.NullString
	Note            sql.NullString
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

func (r *SqlLitePomodoroRepository) findManyPomodoros(where string, args ...interface{}) ([]*pomodoro.Pomodoro, error) {
	query := "SELECT id, duration, status, start_time, finish_time, note FROM pomodoro"
	if where != "" {
		query += " WHERE " + where
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pomodoros []*pomodoro.Pomodoro
	for rows.Next() {
		var dto pomodoroDTO
		err := rows.Scan(&dto.ID, &dto.PlannedDuration, &dto.Status, &dto.StartTime, &dto.FinishTime, &dto.Note)
		if err != nil {
			return nil, err
		}

		startTime, err := time.Parse(time.RFC3339Nano, dto.StartTime)
		if err != nil {
			return nil, err
		}

		var finishTime time.Time
		if dto.FinishTime.Valid {
			finishTime, err = time.Parse(time.RFC3339Nano, dto.FinishTime.String)
			if err != nil {
				return nil, err
			}
		}

		pomodoros = append(pomodoros, &pomodoro.Pomodoro{
			ID:              dto.ID,
			PlannedDuration: time.Duration(dto.PlannedDuration),
			StartTime:       startTime,
			FinishTime:      finishTime,
			Status:          pomodoro.Status(dto.Status),
			Note:            dto.Note.String,
		})
	}

	return pomodoros, nil
}

func (r *SqlLitePomodoroRepository) findPomodoro(where string, args ...interface{}) (*pomodoro.Pomodoro, error) {
	pomodoros, err := r.findManyPomodoros(where, args...)
	if err != nil {
		return nil, err
	}

	if len(pomodoros) == 0 {
		return nil, nil
	}

	// todo: if more than one pomodoro is found, return an error

	return pomodoros[0], nil
}

func (r *SqlLitePomodoroRepository) FindByID(id string) (*pomodoro.Pomodoro, error) {
	return r.findPomodoro("id = ?", id)
}

func (r *SqlLitePomodoroRepository) FindActive() (*pomodoro.Pomodoro, error) {
	return r.findPomodoro("status = ?", pomodoro.StatusRunning)
}

func (r *SqlLitePomodoroRepository) FindAll() ([]*pomodoro.Pomodoro, error) {
	return r.findManyPomodoros("")
}

func (r *SqlLitePomodoroRepository) Update(p *pomodoro.Pomodoro) error {
	dto := pomodoroDTO{
		ID:              p.ID,
		PlannedDuration: int(p.PlannedDuration),
		Status:          string(p.Status),
		StartTime:       p.StartTime.Format(time.RFC3339Nano),
		FinishTime:      sql.NullString{String: p.FinishTime.Format(time.RFC3339Nano), Valid: !p.FinishTime.IsZero()},
		Note:            sql.NullString{String: p.Note, Valid: p.Note != ""},
	}

	_, err := r.db.Exec(
		"UPDATE pomodoro SET duration = ?, status = ?, start_time = ?, finish_time = ?, note = ? WHERE id = ?",
		dto.PlannedDuration, dto.Status, dto.StartTime, dto.FinishTime, dto.Note, dto.ID,
	)
	return err
}

func NewSqlLitePomodoroRepository(db *sql.DB) *SqlLitePomodoroRepository {
	return &SqlLitePomodoroRepository{
		db: db,
	}
}
