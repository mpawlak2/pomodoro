package pomodoro

type PomodoroRepository interface {
	Create(p *Pomodoro) error
	FindByID(id string) (*Pomodoro, error)
	FindAll() ([]*Pomodoro, error)
}
