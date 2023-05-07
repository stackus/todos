package domain

import (
	"github.com/google/uuid"
)

type TodoRepository interface {
	Add(description string) *Todo
	Remove(id uuid.UUID)
	Update(id uuid.UUID, completed bool, description string) *Todo
	Search(search string) []*Todo
	All() []*Todo
	Get(id uuid.UUID) *Todo
	Reorder(ids []uuid.UUID) []*Todo
}
