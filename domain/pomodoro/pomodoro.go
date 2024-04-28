package pomodoro

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "Pending"
	StatusRunning  Status = "Running"
	StatusFinished Status = "Finished"
)

type Pomodoro struct {
	ID       string
	Duration time.Duration
	Status   Status
	Note     string

	startTime time.Time
}

func (p *Pomodoro) Start() {
	p.startTime = time.Now()
	p.Status = StatusRunning
}

func (p *Pomodoro) Finish(note string) {
	p.Note = note
	p.Status = StatusFinished
}

func (p *Pomodoro) RemainingTime() time.Duration {
	if p.Status == StatusRunning {
		elapsed := time.Since(p.startTime)
		return p.Duration - elapsed
	}

	return p.Duration
}

func NewPomodoro(duration time.Duration) *Pomodoro {
	return &Pomodoro{
		ID:       uuid.New().String(),
		Duration: duration,
		Status:   StatusPending,
	}
}
