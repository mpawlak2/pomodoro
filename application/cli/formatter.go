package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/mpawlak2/pomodoro/domain/pomodoro"
)

type PomodoroFormatter struct {
	pomo *pomodoro.Pomodoro
}

func (f *PomodoroFormatter) String() string {
	var builder strings.Builder
	builder.WriteString(color.YellowString(fmt.Sprintf("pomodoro: %s", f.pomo.ID)))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%-20s %s", "Planned duration:", f.pomo.PlannedDuration))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%-20s %s", "Elapsed duration:", f.pomo.ElapsedDuration()))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%-20s %s", "Status:", f.pomo.Status))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%-20s %s", "Start time:", f.pomo.StartTime))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("%-20s %s", "Finish time:", f.pomo.FinishTime))

	builder.WriteString("\n")
	builder.WriteString("\n")
	note := f.pomo.Note
	if note != "" {
		note = strings.TrimSuffix(note, "\n")
	}
	builder.WriteString(fmt.Sprintf("    %s", note))
	builder.WriteString("\n")

	return builder.String()
}

func (f *PomodoroFormatter) RemainingTime() string {
	remainingTime := f.pomo.PlannedDuration - f.pomo.ElapsedDuration()
	minutes := int(remainingTime.Minutes())
	seconds := int(remainingTime.Seconds()) - (minutes * 60)
	return fmt.Sprintf("🍅%02dm%02ds", minutes, seconds)
}

func NewPomodoroFormatter(pomo *pomodoro.Pomodoro) *PomodoroFormatter {
	return &PomodoroFormatter{
		pomo: pomo,
	}
}
