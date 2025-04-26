package main

import (
	"fmt"
	"io/fs"
	"iter"
	"log/slog"
	"os"
	"path"
	"strings"
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	dir := os.DirFS("../../lessons/")

	fmt.Printf("## Exercise Lessons\n\n")

	var currLesson string
	for err, filename := range listExercises(dir) {
		if err != nil {
			return err
		}
		if currLesson != filename.lesson {
			if currLesson != "" {
				fmt.Printf("\n")
			}
			fmt.Printf("### %s\n\n", filename.lesson)
			currLesson = filename.lesson
		}
		fmt.Printf("- [%s](%s)\n", filename.name,
			fmt.Sprintf("<%s>", path.Join("lessons", filename.path, "README.md")),
		)
	}

	return nil
}

type entry struct {
	lesson string
	name   string
	path   string
}

func listExercises(dir fs.FS) iter.Seq2[error, entry] {
	return func(yield func(error, entry) bool) {
		lessons, err := fs.ReadDir(dir, ".")
		if err != nil {
			yield(err, entry{})
			return
		}

		for _, lesson := range lessons {
			pages, err := fs.ReadDir(dir, lesson.Name())
			if err != nil {
				yield(err, entry{})
				return
			}
			for _, page := range pages {
				if strings.Contains(page.Name(), "Exercise: ") {
					if !yield(nil, entry{
						lesson: lesson.Name(),
						name:   page.Name(),
						path:   path.Join(lesson.Name(), page.Name()),
					}) {
						return
					}
				}
			}
		}
	}
}
