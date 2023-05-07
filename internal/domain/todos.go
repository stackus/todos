package domain

import (
	"strings"

	"github.com/google/uuid"
)

// Todos is a list of Todo
type Todos []*Todo

// NewTodos creates a new list of todos
func NewTodos() *Todos {
	return &Todos{}
}

// Add adds a todo to the list
func (l *Todos) Add(description string) *Todo {
	todo := NewTodo(description)
	*l = append(*l, todo)
	return todo
}

// Remove removes a todo from the list
func (l *Todos) Remove(id uuid.UUID) {
	index := l.indexOf(id)
	if index == -1 {
		return
	}
	*l = append((*l)[:index], (*l)[index+1:]...)
}

// Update updates a todo in the list
func (l *Todos) Update(id uuid.UUID, completed bool, description string) *Todo {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}
	todo := (*l)[index]
	todo.Update(completed, description)

	return todo
}

// Search returns a list of todos that match the search string
func (l *Todos) Search(search string) []*Todo {
	list := make([]*Todo, 0)
	for _, todo := range *l {
		if strings.Contains(todo.Description, search) {
			list = append(list, todo)
		}
	}
	return list
}

// All returns a copy of the list of todos
func (l *Todos) All() []*Todo {
	list := make([]*Todo, len(*l))
	copy(list, *l)
	return list
}

// Get returns a todo by id
func (l *Todos) Get(id uuid.UUID) *Todo {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}
	return (*l)[index]
}

// Reorder reorders the list of todos
func (l *Todos) Reorder(ids []uuid.UUID) []*Todo {
	newTodos := make([]*Todo, len(ids))
	for i, id := range ids {
		newTodos[i] = (*l)[l.indexOf(id)]
	}
	copy(*l, newTodos)
	return newTodos
}

// indexOf returns the index of the todo with the given id or -1 if not found
func (l *Todos) indexOf(id uuid.UUID) int {
	for i, todo := range *l {
		if todo.ID == id {
			return i
		}
	}
	return -1
}
