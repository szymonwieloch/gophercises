package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task - a simple CLI task management system",
	Long: `task is a simple CLI task management system

It allows you to define, remove and analyze historical tasks.
It stores the data locally on your machine`,
	Version: "0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing root command with args", args)
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
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
