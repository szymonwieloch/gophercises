package cmd

import (
	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Short:   "Completes the given task",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseTaskID(args[0])
		if err != nil {
			return err
		}
		handleResult(pkg.Complete(id))
		return nil
	},
}
