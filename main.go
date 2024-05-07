package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mpawlak2/pomodoro/application/cli"
)

func main() {
	cli.RunApplication()
}
