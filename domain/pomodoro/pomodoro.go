package pomodoro

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "Pending"
	StatusRunning   Status = "Running"
	StatusFinished  Status = "Finished"
	StatusCancelled Status = "Cancelled"
)

type Pomodoro struct {
	ID              string
	PlannedDuration time.Duration
	Status          Status
	Note            string
	StartTime       time.Time
	FinishTime      time.Time
}

func (p *Pomodoro) Start() {
	p.StartTime = time.Now()
	p.Status = StatusRunning
}

func (p *Pomodoro) Finish(note string) {
	p.Note = note
	p.Status = StatusFinished
	p.FinishTime = time.Now()
}

func (p *Pomodoro) Cancel() {
	p.Status = StatusCancelled
	p.FinishTime = time.Now()
}

func (p *Pomodoro) RemainingTime() time.Duration {
	if p.Status == StatusRunning {
		elapsed := time.Since(p.StartTime)
		return p.PlannedDuration - elapsed
	}

	return p.PlannedDuration
}

func (p *Pomodoro) ElapsedDuration() time.Duration {
	if p.Status == StatusRunning {
		return time.Since(p.StartTime)
	}

	return p.FinishTime.Sub(p.StartTime)
}

func NewPomodoro(plannedDuration time.Duration) *Pomodoro {
	return &Pomodoro{
		ID:              uuid.New().String(),
		PlannedDuration: plannedDuration,
		Status:          StatusPending,
	}
}
