package home

import (
	"context"

	"github.com/stackus/todos/internal/domain"
)

type (
	Service interface {
		// List returns a copy of the todos list
		List(ctx context.Context) ([]*domain.Todo, error)
	}

	service struct {
		todos domain.TodoRepository
	}
)

func NewService(todos domain.TodoRepository) Service {
	return &service{
		todos: todos,
	}
}

func (s service) List(context.Context) ([]*domain.Todo, error) {
	return s.todos.All(), nil
}
