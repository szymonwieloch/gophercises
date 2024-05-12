package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task - a simple CLI task management system",
	Long: `task is a simple CLI task management system

It allows you to define, remove and analyze historical tasks.
It stores the data locally on your machine`,
	Version: "0.1",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		handleResult(pkg.List(true, false))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
