package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/GourangaDasSamrat/todo-cli-go/internal/models"
	"github.com/fatih/color"
)

var (
	// Priority colors
	HighPriorityColor   = color.New(color.FgRed, color.Bold)
	MediumPriorityColor = color.New(color.FgYellow)
	LowPriorityColor    = color.New(color.FgBlue)

	// Status colors
	CompletedColor = color.New(color.FgGreen)
	OverdueColor   = color.New(color.FgRed)
	PendingColor   = color.New(color.FgWhite)

	// UI colors
	HeaderColor  = color.New(color.FgCyan, color.Bold)
	ErrorColor   = color.New(color.FgRed, color.Bold)
	SuccessColor = color.New(color.FgGreen, color.Bold)
	WarningColor = color.New(color.FgYellow)
	InfoColor    = color.New(color.FgCyan)
)

// PrintTask prints a single task with colors
func PrintTask(task *models.Task) {
	// Print ID and Title
	fmt.Printf("ID: %s\n", task.ID)

	// Print title with priority color
	titleColor := getPriorityColor(task.Priority)
	titleColor.Printf("Title: %s\n", task.Title)

	// Print description
	if task.Description != "" {
		fmt.Printf("Description: %s\n", task.Description)
	}

	// Print priority
	fmt.Printf("Priority: ")
	titleColor.Printf("%s\n", task.Priority.String())

	// Print status
	fmt.Printf("Status: ")
	statusColor := getStatusColor(task.Status)
	statusColor.Printf("%s\n", task.Status.String())

	// Print project
	if task.Project != "" {
		fmt.Printf("Project: %s\n", task.Project)
	}

	// Print tags
	if len(task.Tags) > 0 {
		fmt.Printf("Tags: %s\n", strings.Join(task.Tags, ", "))
	}

	// Print due date
	if !task.DueDate.IsZero() {
		fmt.Printf("Due Date: %s", task.DueDate.Format("2006-01-02 15:04"))
		if task.IsOverdue() {
			OverdueColor.Printf(" (OVERDUE)")
		}
		fmt.Println()
	}

	// Print created date
	fmt.Printf("Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04"))

	// Print completed date
	if !task.CompletedAt.IsZero() {
		fmt.Printf("Completed: %s\n", task.CompletedAt.Format("2006-01-02 15:04"))
	}

	fmt.Println(strings.Repeat("-", 60))
}

// PrintTaskList prints a list of tasks in table format
func PrintTaskList(tasks []*models.Task) {
	if len(tasks) == 0 {
		InfoColor.Println("No tasks found.")
		return
	}

	// Print header
	HeaderColor.Println(strings.Repeat("=", 120))
	HeaderColor.Printf("%-8s %-30s %-10s %-12s %-15s %-20s %-15s\n",
		"ID", "Title", "Priority", "Status", "Project", "Due Date", "Tags")
	HeaderColor.Println(strings.Repeat("=", 120))

	// Print tasks
	for _, task := range tasks {
		printTaskRow(task)
	}

	fmt.Println()
	InfoColor.Printf("Total: %d tasks\n", len(tasks))
}

func printTaskRow(task *models.Task) {
	// Truncate title if too long
	title := task.Title
	if len(title) > 28 {
		title = title[:25] + "..."
	}

	// Truncate project if too long
	project := task.Project
	if len(project) > 13 {
		project = project[:10] + "..."
	}

	// Format due date
	dueDate := ""
	if !task.DueDate.IsZero() {
		dueDate = task.DueDate.Format("2006-01-02 15:04")
		if task.IsOverdue() {
			dueDate += " ⚠"
		}
	}

	// Format tags
	tags := ""
	if len(task.Tags) > 0 {
		tags = strings.Join(task.Tags, ", ")
		if len(tags) > 13 {
			tags = tags[:10] + "..."
		}
	}

	// Get colors
	priorityColor := getPriorityColor(task.Priority)
	statusColor := getStatusColor(task.Status)

	// Print row
	fmt.Printf("%-8s ", task.ID)
	priorityColor.Printf("%-30s ", title)
	priorityColor.Printf("%-10s ", task.Priority.String())
	statusColor.Printf("%-12s ", task.Status.String())
	fmt.Printf("%-15s ", project)
	fmt.Printf("%-20s ", dueDate)
	fmt.Printf("%-15s\n", tags)
}

// PrintSuccess prints a success message
func PrintSuccess(message string) {
	SuccessColor.Printf("✓ %s\n", message)
}

// PrintError prints an error message
func PrintError(message string) {
	ErrorColor.Printf("✗ %s\n", message)
}

// PrintWarning prints a warning message
func PrintWarning(message string) {
	WarningColor.Printf("⚠ %s\n", message)
}

// PrintInfo prints an info message
func PrintInfo(message string) {
	InfoColor.Printf("ℹ %s\n", message)
}

// PrintHeader prints a section header
func PrintHeader(message string) {
	HeaderColor.Printf("\n=== %s ===\n\n", message)
}

func getPriorityColor(priority models.Priority) *color.Color {
	switch priority {
	case models.PriorityHigh:
		return HighPriorityColor
	case models.PriorityMedium:
		return MediumPriorityColor
	case models.PriorityLow:
		return LowPriorityColor
	default:
		return color.New(color.FgWhite)
	}
}

func getStatusColor(status models.Status) *color.Color {
	switch status {
	case models.StatusCompleted:
		return CompletedColor
	case models.StatusOverdue:
		return OverdueColor
	case models.StatusPending:
		return PendingColor
	default:
		return color.New(color.FgWhite)
	}
}

// FormatDuration formats a duration in human-readable format
func FormatDuration(d time.Duration) string {
	if d < 0 {
		return "overdue"
	}

	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24

	if days > 0 {
		return fmt.Sprintf("%dd %dh", days, hours)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	}
	return fmt.Sprintf("%dm", int(d.Minutes()))
}
