package cmd

import (
	"github.com/GourangaDasSamrat/todo-cli-go/internal/ui"
	"github.com/GourangaDasSamrat/todo-cli-go/pkg/filter"
	sortpkg "github.com/GourangaDasSamrat/todo-cli-go/pkg/sort"
	"github.com/spf13/cobra"
)

var searchKeyword string

var searchCmd = &cobra.Command{
	Use:     "search [keyword]",
	Aliases: []string{"find", "s"},
	Short:   "Search tasks by keyword",
	Long:    `Search tasks by keyword in title or description.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchKeyword = args[0]

		// Load tasks
		taskList, err := store.Load()
		if err != nil {
			ui.PrintError("Failed to load tasks: " + err.Error())
			return
		}

		// Apply filter
		f := filter.NewKeywordFilter(searchKeyword)
		tasks := f.Apply(taskList.Tasks)

		// Sort by relevance (created date for now)
		sortpkg.Sort(tasks, sortpkg.SortByCreatedAt, false)

		// Display results
		ui.PrintHeader("Search Results for: " + searchKeyword)
		ui.PrintTaskList(tasks)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
