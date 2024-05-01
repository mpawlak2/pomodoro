package pomodoro

import "time"

type ActivePomodoroExistsError struct{}

func (e *ActivePomodoroExistsError) Error() string {
	return "an active pomodoro already exists"
}

type PomodoroService struct {
	repository PomodoroRepository
}

func (s *PomodoroService) CreatePomodoro(plannedDuration time.Duration) (*Pomodoro, error) {
	if active, err := s.repository.FindActive(); err == nil && active != nil {
		return nil, &ActivePomodoroExistsError{}
	}
	p := NewPomodoro(plannedDuration)
	return p, nil
}

func NewPomodoroService(repository PomodoroRepository) *PomodoroService {
	return &PomodoroService{
		repository: repository,
	}
}
