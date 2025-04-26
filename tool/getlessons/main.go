package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	lessons, err := GetLessonFromFile("lessons.json")
	if err != nil {
		return err
	}

	return WriteFiles(lessons, "../../lessons/")
}
