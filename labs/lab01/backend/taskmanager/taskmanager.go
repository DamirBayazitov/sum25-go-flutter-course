package taskmanager

import (
	"errors"
	"time"
)

// Predefined errors
var (
	ErrTaskNotFound = errors.New("task not found")
	ErrEmptyTitle   = errors.New("title cannot be empty")
)

// Task represents a single task
type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

// TaskManager manages a collection of tasks
type TaskManager struct {
	tasks  map[int]Task
	nextID int
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	// TODO: Implement this function
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// AddTask adds a new task to the manager, returns an error if the title is empty, and increments the nextID
func (tm *TaskManager) AddTask(title, description string) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// UpdateTask updates an existing task, returns an error if the title is empty or the task is not found
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	// TODO: Implement this functions
	_, ok := tm.tasks[id]
	if !ok { // if the task with this id does not exist, then we return an error message
		return ErrTaskNotFound
	}
	if title == "" { // if the title is empty, then we return an error message
		return ErrEmptyTitle
	}
	tm.tasks[id].Title = title
	tm.tasks[id].Description = description
	tm.tasks[id].Done = done
	tm.tasks[id].ID = id
	return nil
}

// DeleteTask removes a task from the manager, returns an error if the task is not found
func (tm *TaskManager) DeleteTask(id int) error {
	// TODO: Implement this function
	_, found := tm.tasks[id]
	if !found { // if the task with this id does not exist, then we return an error message
		return ErrTaskNotFound
	}
	delete(tm.tasks, id)
	return nil
}

// GetTask retrieves a task by ID, returns an error if the task is not found
func (tm *TaskManager) GetTask(id int) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// ListTasks returns all tasks, optionally filtered by done status, returns an empty slice if no tasks are found
func (tm *TaskManager) ListTasks(filterDone *bool) []Task {
	// TODO: Implement this function
	var tasks []*Task = []*Task{}
	for _, value := range tm.tasks {
		if value.Done == *filterDone {
			tasks = append(tasks, value)
		}
	}
	return tasks
}
