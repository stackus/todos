package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// NewTodo creates a new todo
func NewTodo(description string) *Todo {
	return &Todo{
		ID:          uuid.New(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

// Update updates a todo
func (t *Todo) Update(completed bool, description string) {
	t.Completed = completed
	t.Description = description
}
