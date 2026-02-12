package cmd

import (
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/spf13/cobra"
)

var (
	completeID       string
	markAsIncomplete bool
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"done", "finish", "c"},
	Short:   "Mark a task as complete or incomplete",
	Long:    `Mark a task as complete or incomplete by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		// Find task
		task := taskList.GetByID(completeID)
		if task == nil {
			ui.PrintError("Task not found with ID: " + completeID)
			return
		}

		// Update status
		if markAsIncomplete {
			task.MarkIncomplete()
			ui.PrintSuccess("Task marked as incomplete!")
		} else {
			task.MarkComplete()
			ui.PrintSuccess("Task marked as complete!")
		}

		// Save
		if err := store.Save(taskList); err != nil {
			ui.PrintError("Failed to save task: " + err.Error())
			return
		}

		ui.PrintTask(task)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	completeCmd.Flags().StringVarP(&completeID, "id", "i", "", "Task ID (required)")
	completeCmd.Flags().BoolVarP(&markAsIncomplete, "incomplete", "u", false, "Mark as incomplete instead")
	completeCmd.MarkFlagRequired("id")
}
