package pomodoro

type PomodoroRepository interface {
	Create(p *Pomodoro) error
	Update(p *Pomodoro) error
	FindActive() (*Pomodoro, error)
	FindByID(id string) (*Pomodoro, error)
	FindAll() ([]*Pomodoro, error)
}
