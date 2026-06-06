package task

import (
	"fmt"
	"time"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Status string

const (
	ToDo       Status = "to_do"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Status	  Status    `json:"status"`
	Priority    Priority  `json:"priority"`
	Tags		[]string  `json:"tags"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ParsePriority(priority string) (Priority, error) {
	switch priority {
		case "h", "high":
			return High, nil
		case "m", "medium":
			return Medium, nil
		case "l", "low":
			return Low, nil
		default:
			return "", fmt.Errorf("invalid priority: %s - use h/m/l or high/medium/low", priority)
	}
}