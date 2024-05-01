package pomodoro

import (
	"testing"
	"time"
)

type MockPomodoroRepository struct {
	pomodoros []*Pomodoro
}

func (r *MockPomodoroRepository) Create(p *Pomodoro) error {
	r.pomodoros = append(r.pomodoros, p)
	return nil
}

func (r *MockPomodoroRepository) FindByID(id string) (*Pomodoro, error) {
	for _, p := range r.pomodoros {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, nil
}

func (r *MockPomodoroRepository) FindAll() ([]*Pomodoro, error) {
	return r.pomodoros, nil
}

func (r *MockPomodoroRepository) FindActive() (*Pomodoro, error) {
	for _, p := range r.pomodoros {
		if p.Status == StatusRunning {
			return p, nil
		}
	}

	return nil, nil
}

func TestShouldNotCreateSecondPomodoroIfAnotherIsActive(t *testing.T) {
	repo := &MockPomodoroRepository{}
	service := NewPomodoroService(repo)
	if service == nil {
		t.Errorf("Expected service to be created, but got nil")
	}

	pomo, err := service.CreatePomodoro(25 * time.Minute)
	if err != nil {
		t.Errorf("Error creating pomodoro: %v", err)
	}
	pomo.Start()
	repo.Create(pomo)

	_, err = service.CreatePomodoro(25 * time.Minute)
	// check if err is of type ActivePomodoroExistsError
	if _, ok := err.(*ActivePomodoroExistsError); !ok {
		t.Errorf("Expected ActivePomodoroExistsError, but got %v", err)
	}
}
