package pomodoro

import (
	"testing"
	"time"
)

func TestPomodoroAggregateRoot(t *testing.T) {
	p := NewPomodoro(25 * time.Minute)

	if p.ID == "" {
		t.Errorf("Expected ID to be set, but got %v", p.ID)
	}

	if p.PlannedDuration != 25*time.Minute {
		t.Errorf("Expected planned duration to be 25 minutes, but got %v", p.PlannedDuration)
	}

	if p.Status != StatusPending {
		t.Errorf("Expected status to be Pending, but got %v", p.Status)
	}

	p.Start()
	if p.Status != StatusRunning {
		t.Errorf("Expected status to be Running, but got %v", p.Status)
	}

	p.Finish("Testing")
	if p.Status != StatusFinished {
		t.Errorf("Expected status to be Finished, but got %v", p.Status)
	}
}

func TestStopAndResumePomodoro(t *testing.T) {
	p := NewPomodoro(25 * time.Minute)

	if p.RemainingTime() != 25*time.Minute {
		t.Errorf("Expected remaining time to be 25 minutes, but got %v", p.RemainingTime())
	}

	p.Start()
	if p.Status != StatusRunning {
		t.Errorf("Expected status to be Running, but got %v", p.Status)
	}

	if p.RemainingTime() == 25*time.Minute {
		t.Errorf("Expected remaining time to be less than 25 minutes once started, but got %v", p.RemainingTime())
	}
}

func TestCompletePomodoro(t *testing.T) {
	p := NewPomodoro(25 * time.Minute)

	p.Start()
	p.Finish("This is a test note.")

	if p.Status != StatusFinished {
		t.Errorf("Expected status to be Finished, but got %v", p.Status)
	}

	if p.FinishTime.IsZero() {
		t.Errorf("Expected finish time to be set, but got zero")
	}

	if p.Note != "This is a test note." {
		t.Errorf("Expected note to be 'This is a test note.', but got %v", p.Note)
	}
}

func TestCancelPomodoro(t *testing.T) {
	p := NewPomodoro(25 * time.Minute)

	p.Start()
	if p.StartTime == (time.Time{}) {
		t.Errorf("Expected start time to be set, but got %v", p.StartTime)
	}
	if p.FinishTime != (time.Time{}) {
		t.Errorf("Expected finish time to be zero, but got %v", p.FinishTime)
	}

	p.Cancel()
	if p.Status != StatusCancelled {
		t.Errorf("Expected status to be Cancelled, but got %v", p.Status)
	}
	if p.FinishTime == (time.Time{}) {
		t.Errorf("Expected finish time to be set, but got %v", p.FinishTime)
	}

	if p.ElapsedDuration() == 0 {
		t.Errorf("Expected elapsed duration to be greater than zero, but got %v", p.ElapsedDuration())
	}
}
