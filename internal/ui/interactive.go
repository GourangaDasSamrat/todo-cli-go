package ui

import (
	"fmt"
	"strings"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/manifoldco/promptui"
)

// InteractiveMenu represents the main interactive menu
type InteractiveMenu struct {
	options []string
}

// NewInteractiveMenu creates a new interactive menu
func NewInteractiveMenu() *InteractiveMenu {
	return &InteractiveMenu{
		options: []string{
			"View All Tasks",
			"Add New Task",
			"Edit Task",
			"Delete Task",
			"Mark Complete/Incomplete",
			"Filter Tasks",
			"Search Tasks",
			"Backup Data",
			"Restore Data",
			"Exit",
		},
	}
}

// Show displays the interactive menu and returns selected option
func (m *InteractiveMenu) Show() (int, error) {
	prompt := promptui.Select{
		Label: "Todo CLI - Main Menu",
		Items: m.options,
		Size:  10,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	return index, nil
}

// PromptTaskInput prompts for task input
func PromptTaskInput() (*models.Task, error) {
	task := &models.Task{}

	// Title
	titlePrompt := promptui.Prompt{
		Label: "Title",
		Validate: func(input string) error {
			if strings.TrimSpace(input) == "" {
				return fmt.Errorf("title cannot be empty")
			}
			return nil
		},
	}
	title, err := titlePrompt.Run()
	if err != nil {
		return nil, err
	}
	task.Title = title

	// Description
	descPrompt := promptui.Prompt{
		Label: "Description (optional)",
	}
	desc, _ := descPrompt.Run()
	task.Description = desc

	// Priority
	priorityPrompt := promptui.Select{
		Label: "Priority",
		Items: []string{"low", "medium", "high"},
	}
	_, priority, err := priorityPrompt.Run()
	if err != nil {
		return nil, err
	}
	task.Priority = models.ParsePriority(priority)

	// Project
	projectPrompt := promptui.Prompt{
		Label: "Project (optional)",
	}
	project, _ := projectPrompt.Run()
	task.Project = project

	// Tags
	tagsPrompt := promptui.Prompt{
		Label: "Tags (comma-separated, optional)",
	}
	tags, _ := tagsPrompt.Run()
	if tags != "" {
		task.Tags = strings.Split(tags, ",")
		for i := range task.Tags {
			task.Tags[i] = strings.TrimSpace(task.Tags[i])
		}
	}

	// Due date
	dueDatePrompt := promptui.Prompt{
		Label: "Due Date (YYYY-MM-DD HH:MM, optional)",
	}
	dueDate, _ := dueDatePrompt.Run()
	if dueDate != "" {
		// Parsing handled by caller
		task.DueDate.UnmarshalText([]byte(dueDate))
	}

	return task, nil
}

// SelectTask prompts user to select a task from a list
func SelectTask(tasks []*models.Task) (*models.Task, error) {
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no tasks available")
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▸ {{ .Title | cyan }} ({{ .Priority | yellow }}) [{{ .Status | green }}]",
		Inactive: "  {{ .Title }} ({{ .Priority }}) [{{ .Status }}]",
		Selected: "✔ {{ .Title | green }}",
	}

	prompt := promptui.Select{
		Label:     "Select Task",
		Items:     tasks,
		Templates: templates,
		Size:      10,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return tasks[index], nil
}

// ConfirmAction prompts for confirmation
func ConfirmAction(message string) bool {
	prompt := promptui.Prompt{
		Label:     message,
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false
	}

	return result == "y" || result == "Y"
}

// PromptInput prompts for simple text input
func PromptInput(label string, required bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	if required {
		prompt.Validate = func(input string) error {
			if strings.TrimSpace(input) == "" {
				return fmt.Errorf("input cannot be empty")
			}
			return nil
		}
	}

	return prompt.Run()
}

// SelectOption prompts user to select from options
func SelectOption(label string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := prompt.Run()
	return result, err
}
