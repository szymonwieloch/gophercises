package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Adds a task to the system",
	RunE: func(cmd *cobra.Command, args []string) error {
		taskName := strings.Join(args, " ")
		taskName = strings.TrimSpace(taskName)
		if taskName == "" {
			return fmt.Errorf("Empty task name")
		}
		return pkg.Add(taskName)
	},
}
