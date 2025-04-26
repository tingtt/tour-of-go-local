package main

import "iter"

type Lessons struct {
	Welcome     Lesson `json:"welcome"`
	Basics      Lesson `json:"basics"`
	FlowControl Lesson `json:"flowcontrol"`
	MoreTypes   Lesson `json:"moretypes"`
	Methods     Lesson `json:"methods"`
	Concurrency Lesson `json:"concurrency"`
}

type IndexedLesson struct {
	Key string
	Lesson
}

func (l Lessons) All() iter.Seq2[int, IndexedLesson] {
	s := []IndexedLesson{
		{Key: "welcome", Lesson: l.Welcome},
		{Key: "basics", Lesson: l.Basics},
		{Key: "flowcontrol", Lesson: l.FlowControl},
		{Key: "moretypes", Lesson: l.MoreTypes},
		{Key: "methods", Lesson: l.Methods},
		{Key: "concurrency", Lesson: l.Concurrency},
	}
	return func(yield func(int, IndexedLesson) bool) {
		for i, lesson := range s {
			if !yield(i, lesson) {
				return
			}
		}
	}
}

type Lesson struct {
	Title       string
	Description string
	Pages       []LessonPage
}

type LessonPage struct {
	Title   string
	Content string
	Files   []LessonPageFile
}

type LessonPageFile struct {
	Name    string
	Content string
	Hash    string
}
