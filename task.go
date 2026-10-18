package main

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var ErrEmptyInput = errors.New("text cannot be empty")

type TaskState string

const (
	StatePending    TaskState = "Pending"
	StateProcessing TaskState = "Processing"
	StateCompleted  TaskState = "Completed"
)

type Result struct {
	Words      int
	Characters int
	Lines      int
}

type Task struct {
	ID     int
	Input  string
	State  TaskState
	Result *Result
}

type Processor interface {
	Process(input string) (Result, error)
}

type TextAnalyzer struct{}

func (a TextAnalyzer) Process(input string) (Result, error) {
	if strings.TrimSpace(input) == "" {
		return Result{}, ErrEmptyInput
	}

	result := Result{
		Words:      len(strings.Fields(input)),
		Characters: utf8.RuneCountInString(input),
		Lines:      strings.Count(input, "\n") + 1,
	}

	return result, nil
}
