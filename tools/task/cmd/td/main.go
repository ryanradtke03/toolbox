package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "td",
		Short: "A simple task management CLI tool",
		Long:  "Task is a command-line application for managing your tasks with priorities, tags, and categories.",

		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}