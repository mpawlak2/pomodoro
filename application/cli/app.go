package cli

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mpawlak2/pomodoro/domain/pomodoro"
	"github.com/mpawlak2/pomodoro/infrastructure/repository"
	"github.com/urfave/cli/v2"
)

func RunApplication() {
	db, err := sql.Open("sqlite3", "pomodoro.db")
	if err != nil {
		panic(err)
	}
	repository.InitializeSqlLiteDB(db)
	defer db.Close()

	repo := repository.NewSqlLitePomodoroRepository(db)
	pomodoroService := pomodoro.NewPomodoroService(repo)

	app := &cli.App{
		Name:  "pd",
		Usage: "pomodoro management system",
		Commands: []*cli.Command{
			{
				Name:    "log",
				Aliases: []string{"a"},
				Usage:   "log all pomodoros",
				Action: func(cCtx *cli.Context) error {
					pomos, err := repo.FindAll()
					if err != nil {
						return err
					}
					for _, pomo := range pomos {
						formatter := NewPomodoroFormatter(pomo)
						fmt.Println(formatter)
					}
					return nil
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "start a new pomodoro",
				Action: func(cCtx *cli.Context) error {
					pomo, err := pomodoroService.CreatePomodoro(25 * time.Minute)
					if err != nil {
						return err
					}
					pomo.Start()
					repo.Create(pomo)
					fmt.Println(pomo)
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "show status of the current pomodoro",
				Action: func(cCtx *cli.Context) error {
					pomo, err := repo.FindActive()
					if err != nil {
						return err
					}

					if pomo == nil {
						return fmt.Errorf("no active pomodoro")
					}
					formatter := NewPomodoroFormatter(pomo)
					fmt.Println(formatter.RemainingTime())
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "finish a pomodoro",
				Action: func(cCtx *cli.Context) error {
					note := cCtx.Args().First()
					if note == "" {
						note, err = GetInputFromTextEditor()
						if err != nil {
							return err
						}
						if note == "" {
							return fmt.Errorf("note is required")
						}
					}
					pomo, err := repo.FindActive()
					if err != nil {
						return err
					}
					pomo.Finish(note)
					repo.Update(pomo)
					fmt.Println(NewPomodoroFormatter(pomo))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
