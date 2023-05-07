package todos

import (
	"context"

	"github.com/google/uuid"

	"github.com/stackus/todos/internal/domain"
)

type (
	Service interface {
		// Add adds a todo to the list
		Add(ctx context.Context, description string) (*domain.Todo, error)
		// Remove removes a todo from the list
		Remove(ctx context.Context, id uuid.UUID) error
		// Update updates a todo in the list
		Update(ctx context.Context, id uuid.UUID, completed bool, description string) (*domain.Todo, error)
		// Search returns a list of todos that match the search string
		Search(ctx context.Context, search string) ([]*domain.Todo, error)
		// Get returns a todo by id
		Get(ctx context.Context, id uuid.UUID) (*domain.Todo, error)
		// Sort sorts the todos by the given ids
		Sort(ctx context.Context, ids []uuid.UUID) error
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

func (s service) Add(_ context.Context, description string) (*domain.Todo, error) {
	todo := s.todos.Add(description)

	return todo, nil
}

func (s service) Remove(_ context.Context, id uuid.UUID) error {
	s.todos.Remove(id)

	return nil
}

func (s service) Update(_ context.Context, id uuid.UUID, completed bool, description string) (*domain.Todo, error) {
	todo := s.todos.Update(id, completed, description)

	return todo, nil
}

func (s service) Search(_ context.Context, search string) ([]*domain.Todo, error) {
	todos := s.todos.Search(search)

	return todos, nil
}

func (s service) Get(_ context.Context, id uuid.UUID) (*domain.Todo, error) {
	todo := s.todos.Get(id)

	return todo, nil
}

func (s service) Sort(_ context.Context, ids []uuid.UUID) error {
	s.todos.Reorder(ids)

	return nil
}
