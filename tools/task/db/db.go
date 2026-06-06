package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryanradtke03/task/internal/task"
)

func dbPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".local", "share", "task", ".tasks.json"), nil
}

func ReadTasks() ([]task.Task, error) {
	path, err := dbPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil // No tasks file, return empty list
		}
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse tasks file: %w", err)
	}

	return tasks, nil
}

func WriteTasks(tasks []task.Task) error {
	path, err := dbPath()
	if err != nil {
		return err
	}
	
	// Crete parent directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create tasks directory: %w", err)
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize tasks: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	return nil
}