package main

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidState = errors.New("task is not pending")
)

type Service struct {
	tasks     map[int]*Task
	nextID    int
	processor Processor
}

func NewService(processor Processor) *Service {
	return &Service{
		tasks:     make(map[int]*Task),
		nextID:    1,
		processor: processor,
	}
}

func (s *Service) CreateTask(input string) (*Task, error) {
	if _, err := s.processor.Process(input); err != nil {
		return nil, err
	}

	task := &Task{
		ID:    s.nextID,
		Input: input,
		State: StatePending,
	}

	s.tasks[task.ID] = task
	s.nextID++

	return task, nil
}

func (s *Service) GetTask(id int) (*Task, error) {
	task, exists := s.tasks[id]

	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (s *Service) ProcessTask(id int) error {
	task, err := s.GetTask(id)
	if err != nil {
		return err
	}

	if task.State != StatePending {
		return ErrInvalidState
	}

	task.State = StateProcessing

	result, err := s.processor.Process(task.Input)
	if err != nil {
		task.State = StatePending
		return err
	}

	task.Result = &result
	task.State = StateCompleted

	return nil
}
