package filter

import (
	"strings"
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
)

// Filter represents a task filter
type Filter struct {
	Status   *models.Status
	Priority *models.Priority
	Project  string
	Tags     []string
	Keyword  string
	DateFrom time.Time
	DateTo   time.Time
}

// Apply applies the filter to a task list
func (f *Filter) Apply(tasks []*models.Task) []*models.Task {
	var filtered []*models.Task

	for _, task := range tasks {
		if f.matches(task) {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func (f *Filter) matches(task *models.Task) bool {
	// Filter by status
	if f.Status != nil && task.Status != *f.Status {
		return false
	}

	// Filter by priority
	if f.Priority != nil && task.Priority != *f.Priority {
		return false
	}

	// Filter by project
	if f.Project != "" && !strings.EqualFold(task.Project, f.Project) {
		return false
	}

	// Filter by tags
	if len(f.Tags) > 0 {
		hasTag := false
		for _, filterTag := range f.Tags {
			for _, taskTag := range task.Tags {
				if strings.EqualFold(taskTag, filterTag) {
					hasTag = true
					break
				}
			}
			if hasTag {
				break
			}
		}
		if !hasTag {
			return false
		}
	}

	// Filter by keyword
	if f.Keyword != "" {
		keyword := strings.ToLower(f.Keyword)
		if !strings.Contains(strings.ToLower(task.Title), keyword) &&
			!strings.Contains(strings.ToLower(task.Description), keyword) {
			return false
		}
	}

	// Filter by date range
	if !f.DateFrom.IsZero() && task.DueDate.Before(f.DateFrom) {
		return false
	}

	if !f.DateTo.IsZero() && task.DueDate.After(f.DateTo) {
		return false
	}

	return true
}

// NewStatusFilter creates a filter for specific status
func NewStatusFilter(status models.Status) *Filter {
	return &Filter{Status: &status}
}

// NewPriorityFilter creates a filter for specific priority
func NewPriorityFilter(priority models.Priority) *Filter {
	return &Filter{Priority: &priority}
}

// NewProjectFilter creates a filter for specific project
func NewProjectFilter(project string) *Filter {
	return &Filter{Project: project}
}

// NewTagFilter creates a filter for specific tags
func NewTagFilter(tags []string) *Filter {
	return &Filter{Tags: tags}
}

// NewKeywordFilter creates a filter for keyword search
func NewKeywordFilter(keyword string) *Filter {
	return &Filter{Keyword: keyword}
}
