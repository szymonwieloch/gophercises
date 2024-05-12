package cmd

import (
	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Short:   "Removes a task from the system",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := parseTaskID(args[0])
		if err != nil {
			return err
		}
		handleResult(pkg.Rm(id))
		return nil
	},
}
