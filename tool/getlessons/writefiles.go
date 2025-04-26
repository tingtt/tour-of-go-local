package main

import (
	"fmt"
	"os"
	"path"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func WriteFiles(lessons Lessons, dest string) error {
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	for i, lesson := range lessons.All() {
		lessonDir := path.Join(dest, fmt.Sprintf("%d. %s", i, lesson.Key))
		err := os.MkdirAll(lessonDir, os.ModePerm)
		if err != nil {
			return err
		}

		readmeBytes := "# " + lesson.Title + "\n\n" + lesson.Description
		err = os.WriteFile(path.Join(lessonDir, "README.md"), []byte(readmeBytes), os.ModePerm)
		if err != nil {
			return err
		}

		for i, page := range lesson.Pages {
			pageDir := path.Join(lessonDir, fmt.Sprintf("%d. %s", i, page.Title))
			err := os.MkdirAll(pageDir, os.ModePerm)
			if err != nil {
				return err
			}

			markdown, err := htmltomarkdown.ConvertString(page.Content)
			if err != nil {
				return err
			}
			err = os.WriteFile(path.Join(pageDir, "README.md"), []byte(markdown), os.ModePerm)
			if err != nil {
				return err
			}

			for _, file := range page.Files {
				err = os.WriteFile(path.Join(pageDir, file.Name), []byte(file.Content), os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
