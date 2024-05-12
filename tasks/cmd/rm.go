package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"r"},
	Short:   "Removes a task from the system",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		pkg.Add(taskName)
	},
}
