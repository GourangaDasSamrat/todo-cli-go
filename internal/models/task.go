package models

import (
	"time"
)

// Priority levels for tasks
type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityLow:
		return "low"
	case PriorityMedium:
		return "medium"
	case PriorityHigh:
		return "high"
	default:
		return "unknown"
	}
}

// ParsePriority converts string to Priority
func ParsePriority(s string) Priority {
	switch s {
	case "low":
		return PriorityLow
	case "medium":
		return PriorityMedium
	case "high":
		return PriorityHigh
	default:
		return PriorityLow
	}
}

// Status represents task completion status
type Status int

const (
	StatusPending Status = iota
	StatusCompleted
	StatusOverdue
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusCompleted:
		return "completed"
	case StatusOverdue:
		return "overdue"
	default:
		return "unknown"
	}
}

// Task represents a todo item
type Task struct {
	ID          string    `json:"id" yaml:"id"`
	Title       string    `json:"title" yaml:"title"`
	Description string    `json:"description" yaml:"description"`
	Priority    Priority  `json:"priority" yaml:"priority"`
	Status      Status    `json:"status" yaml:"status"`
	Tags        []string  `json:"tags" yaml:"tags"`
	Project     string    `json:"project" yaml:"project"`
	DueDate     time.Time `json:"due_date" yaml:"due_date"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" yaml:"updated_at"`
	CompletedAt time.Time `json:"completed_at,omitempty" yaml:"completed_at,omitempty"`
}

// IsOverdue checks if task is overdue
func (t *Task) IsOverdue() bool {
	if t.Status == StatusCompleted {
		return false
	}
	return !t.DueDate.IsZero() && time.Now().After(t.DueDate)
}

// UpdateStatus updates task status and handles overdue logic
func (t *Task) UpdateStatus() {
	if t.Status != StatusCompleted && t.IsOverdue() {
		t.Status = StatusOverdue
	}
}

// MarkComplete marks task as completed
func (t *Task) MarkComplete() {
	t.Status = StatusCompleted
	t.CompletedAt = time.Now()
	t.UpdatedAt = time.Now()
}

// MarkIncomplete marks task as incomplete
func (t *Task) MarkIncomplete() {
	t.Status = StatusPending
	t.CompletedAt = time.Time{}
	t.UpdatedAt = time.Now()
	t.UpdateStatus()
}

// TaskList represents a collection of tasks
type TaskList struct {
	Tasks []*Task `json:"tasks" yaml:"tasks"`
}

// Add adds a task to the list
func (tl *TaskList) Add(task *Task) {
	tl.Tasks = append(tl.Tasks, task)
}

// Remove removes a task by ID
func (tl *TaskList) Remove(id string) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return true
		}
	}
	return false
}

// GetByID retrieves a task by ID
func (tl *TaskList) GetByID(id string) *Task {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return task
		}
	}
	return nil
}

// UpdateAllStatuses updates status for all tasks
func (tl *TaskList) UpdateAllStatuses() {
	for _, task := range tl.Tasks {
		task.UpdateStatus()
	}
}
