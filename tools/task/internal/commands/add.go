package commands


import (
	"fmt"
	"time"

	"github.com/matoous/go-nanoid/v2"
	"github.com/spf13/cobra"
	"github.com/ryanradtke03/td/internal/db"
	"github.com/ryanradtke03/td/internal/task"
)

func AddCommand() *cobra.Command {
	var priority string
	var tags []string
	var category string
	var description string

	cmd := &cobra.Command{
		Use:   "add <title>",
		Short: "Add a new task",
		Long:  "Add a new task with optional priority, tags, category, and description.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			title := args[0]

			p, err := task.ParsePriority(priority)
			if err != nil {
				return err
			}

			// Build the new task
			newTask := task.Task{
				ID:          nanoid.Must(8),
				Title:       title,
				Status:      task.ToDo,
				Priority:    p,
				Tags:        tags,
				Category:    category,
				Description: description,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			// Read existing tasks
			tasks, err := db.ReadTasks()
			if err != nil {
				return fmt.Errorf("failed to read tasks: %w", err)
			}

			// Append the new task and write back to the database
			tasks = append(tasks, newTask)
			if err := db.WriteTasks(tasks); err != nil {
				return fmt.Errorf("failed to write tasks: %w", err)
			}
			
			fmt.Printf("Task added: %s\n", newTask.Title)
			return nil
		},
	}
	cmd.Flags().StringVarP(&priority, "priority", "p", "medium", "Set task priority (h/m/l or high/medium/low)")
	cmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Comma-separated list of tags")
	cmd.Flags().StringVarP(&category, "category", "c", "", "Set task category")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Set task description")

	return cmd
}