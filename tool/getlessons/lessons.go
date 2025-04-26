package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetLessonFromFile(filename string) (Lessons, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Lessons{}, err
	}
	defer file.Close()

	var lessons Lessons
	if err := json.NewDecoder(file).Decode(&lessons); err != nil {
		return Lessons{}, fmt.Errorf("failed to decode lesson: %w", err)
	}
	return lessons, nil
}
