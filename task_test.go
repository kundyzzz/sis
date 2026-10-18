package main

import (
	"errors"
	"testing"
)

func newTestService() *Service {
	return NewService(TextAnalyzer{})
}

func TestCreateTask(t *testing.T) {
	service := newTestService()

	task, err := service.CreateTask("Hello Go world")
	if err != nil {
		t.Fatal(err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.State != StatePending {
		t.Errorf("expected Pending, got %s", task.State)
	}
}

func TestInvalidInput(t *testing.T) {
	service := newTestService()
	inputs := []string{"", "   ", "\n\t"}

	for _, input := range inputs {
		_, err := service.CreateTask(input)
		if !errors.Is(err, ErrEmptyInput) {
			t.Errorf("expected ErrEmptyInput for %q, got %v", input, err)
		}
	}

	if len(service.tasks) != 0 {
		t.Error("invalid tasks must not be stored")
	}
}

func TestGetTask(t *testing.T) {
	service := newTestService()

	created, err := service.CreateTask("Hello")
	if err != nil {
		t.Fatal(err)
	}

	found, err := service.GetTask(created.ID)
	if err != nil {
		t.Fatal(err)
	}

	if found.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, found.ID)
	}
}

func TestGetMissingTask(t *testing.T) {
	service := newTestService()

	_, err := service.GetTask(999)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestProcessTask(t *testing.T) {
	service := newTestService()

	task, err := service.CreateTask("Hello Go world")
	if err != nil {
		t.Fatal(err)
	}

	if err := service.ProcessTask(task.ID); err != nil {
		t.Fatal(err)
	}

	if task.State != StateCompleted {
		t.Errorf("expected Completed, got %s", task.State)
	}

	if task.Result == nil {
		t.Fatal("expected a result")
	}

	if task.Result.Words != 3 {
		t.Errorf("expected 3 words, got %d", task.Result.Words)
	}

	if task.Result.Characters != 14 {
		t.Errorf("expected 14 characters, got %d", task.Result.Characters)
	}

	if task.Result.Lines != 1 {
		t.Errorf("expected 1 line, got %d", task.Result.Lines)
	}
}

func TestProcessTaskTwice(t *testing.T) {
	service := newTestService()

	task, err := service.CreateTask("Hello")
	if err != nil {
		t.Fatal(err)
	}

	if err := service.ProcessTask(task.ID); err != nil {
		t.Fatal(err)
	}

	err = service.ProcessTask(task.ID)
	if !errors.Is(err, ErrInvalidState) {
		t.Errorf("expected ErrInvalidState, got %v", err)
	}
}

func TestIndependentTasks(t *testing.T) {
	service := newTestService()

	first, err := service.CreateTask("Hello Go")
	if err != nil {
		t.Fatal(err)
	}

	second, err := service.CreateTask("One two three")
	if err != nil {
		t.Fatal(err)
	}

	if first.ID == second.ID {
		t.Fatal("task IDs must be unique")
	}

	if err := service.ProcessTask(first.ID); err != nil {
		t.Fatal(err)
	}

	if first.Result.Words != 2 {
		t.Errorf("expected 2 words, got %d", first.Result.Words)
	}

	if second.Result != nil {
		t.Error("second task must not have a result yet")
	}

	if second.State != StatePending {
		t.Errorf("expected second task to remain Pending, got %s", second.State)
	}
}

func TestProcessTextTableDriven(t *testing.T) {
	analyzer := TextAnalyzer{}

	tests := []struct {
		name       string
		input      string
		words      int
		characters int
		lines      int
	}{
		{
			name:       "simple text",
			input:      "Hello Go",
			words:      2,
			characters: 8,
			lines:      1,
		},
		{
			name:       "multiple lines",
			input:      "Go\nLang",
			words:      2,
			characters: 7,
			lines:      2,
		},
		{
			name:       "unicode text",
			input:      "Привет",
			words:      1,
			characters: 6,
			lines:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := analyzer.Process(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			if result.Words != tt.words ||
				result.Characters != tt.characters ||
				result.Lines != tt.lines {
				t.Errorf("unexpected result: %+v", result)
			}
		})
	}
}
