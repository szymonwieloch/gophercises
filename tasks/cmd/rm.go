package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Short:   "Removes a task from the system",
	RunE: func(cmd *cobra.Command, args []string) error {
		number, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return err
		}
		return pkg.Rm(pkg.TaskID(number))
	},
}
