package sort

import (
	"sort"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
)

// SortBy represents the sorting criteria
type SortBy int

const (
	SortByPriority SortBy = iota
	SortByDueDate
	SortByCreatedAt
	SortByTitle
)

// Sort sorts tasks based on the given criteria
func Sort(tasks []*models.Task, by SortBy, ascending bool) {
	switch by {
	case SortByPriority:
		sortByPriority(tasks, ascending)
	case SortByDueDate:
		sortByDueDate(tasks, ascending)
	case SortByCreatedAt:
		sortByCreatedAt(tasks, ascending)
	case SortByTitle:
		sortByTitle(tasks, ascending)
	}
}

func sortByPriority(tasks []*models.Task, ascending bool) {
	sort.Slice(tasks, func(i, j int) bool {
		if ascending {
			return tasks[i].Priority < tasks[j].Priority
		}
		return tasks[i].Priority > tasks[j].Priority
	})
}

func sortByDueDate(tasks []*models.Task, ascending bool) {
	sort.Slice(tasks, func(i, j int) bool {
		// Handle zero times (no due date)
		if tasks[i].DueDate.IsZero() && tasks[j].DueDate.IsZero() {
			return false
		}
		if tasks[i].DueDate.IsZero() {
			return false
		}
		if tasks[j].DueDate.IsZero() {
			return true
		}

		if ascending {
			return tasks[i].DueDate.Before(tasks[j].DueDate)
		}
		return tasks[i].DueDate.After(tasks[j].DueDate)
	})
}

func sortByCreatedAt(tasks []*models.Task, ascending bool) {
	sort.Slice(tasks, func(i, j int) bool {
		if ascending {
			return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)
		}
		return tasks[i].CreatedAt.After(tasks[j].CreatedAt)
	})
}

func sortByTitle(tasks []*models.Task, ascending bool) {
	sort.Slice(tasks, func(i, j int) bool {
		if ascending {
			return tasks[i].Title < tasks[j].Title
		}
		return tasks[i].Title > tasks[j].Title
	})
}
