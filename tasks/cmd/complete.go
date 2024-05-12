package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Short:   "Completes the given task",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		pkg.Add(taskName)
	},
}
