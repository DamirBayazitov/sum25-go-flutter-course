package taskmanager

import (
	"errors"
	"time"
)

var (
	// ErrTaskNotFound is returned when a task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrEmptyTitle is returned when the task title is empty
	ErrEmptyTitle = errors.New("task title cannot be empty")
	// ErrInvalidID is returned when the task ID is invalid
	ErrInvalidID = errors.New("invalid task ID")
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
	tasks  map[int]*Task
	nextID int
}

// NewTaskManager creates a new task manager
func NewTaskManager() *TaskManager {
	// TODO: Implement task manager initialization
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

// AddTask adds a new task to the manager
func (tm *TaskManager) AddTask(title, description string) (*Task, error) {
	if title == "" { // if a title is empty, then we return the error message
		return nil, ErrEmptyTitle
	}
	task := &Task{ // creating an instance of the type Task
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
	tm.tasks[tm.nextID] = task
	tm.nextID++
	return task, nil
}

// UpdateTask updates an existing task
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	// TODO: Implement task updates
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

// DeleteTask removes a task from the manager
func (tm *TaskManager) DeleteTask(id int) error {
	// TODO: Implement task deletion
	_, found := tm.tasks[id]
	if !found { // if the task with this id does not exist, then we return an error message
		return ErrTaskNotFound
	}
	delete(tm.tasks, id)
	return nil
}

// GetTask retrieves a task by ID
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	task, found := tm.tasks[id]
	if !found { // if the task with this id does not exist, then we return an error message
		return nil, ErrTaskNotFound
	}
	return task, nil
}

// ListTasks returns all tasks, optionally filtered by done status
func (tm *TaskManager) ListTasks(filterDone *bool) []*Task {
	// TODO: Implement task listing with optional filter
	var tasks []*Task = []*Task{}
	for _, value := range tm.tasks {
		if value.Done == *filterDone {
			tasks = append(tasks, value)
		}
	}
	return tasks
}
