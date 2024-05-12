package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Adds a task to the system",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		pkg.Add(taskName)
	},
}
