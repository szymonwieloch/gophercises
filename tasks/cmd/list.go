package cmd

import (
	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var (
	listCompleted bool
	listOpen      bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List tasks in the system",
	RunE: func(cmd *cobra.Command, args []string) error {
		pkg.List(listOpen, listCompleted)
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listCompleted, "completed", "c", false, "List completed tasks")
	listCmd.Flags().BoolVarP(&listOpen, "open", "o", false, "List open tasks")
}
