package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewTodo(t *testing.T) {
	type args struct {
		description string
	}
	tests := map[string]struct {
		args args
		want *Todo
	}{
		"NewTodo": {
			args: args{
				description: "test",
			},
			want: &Todo{
				ID:          uuid.New(),
				Description: "test",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Act
			got := NewTodo(tt.args.description)
			// Assert
			if got.Description != tt.want.Description {
				t.Errorf("NewTodo() = %v, want %v", got.Description, tt.want.Description)
			}
			if got.Completed != tt.want.Completed {
				t.Errorf("NewTodo() = %v, want %v", got.Completed, tt.want.Completed)
			}
		})
	}
}

func TestTodo_Update(t *testing.T) {
	type fields struct {
		ID          uuid.UUID
		Description string
		Completed   bool
		CreatedAt   time.Time
	}
	type args struct {
		completed   bool
		description string
	}
	tests := map[string]struct {
		fields fields
		args   args
	}{
		"Update": {
			fields: fields{
				ID:          uuid.New(),
				Description: "test",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
			args: args{
				completed:   true,
				description: "test2",
			},
		},
		"SameValues:": {
			fields: fields{
				ID:          uuid.New(),
				Description: "test",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
			args: args{
				completed:   false,
				description: "test",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t1 *testing.T) {
			todo := &Todo{
				ID:          tt.fields.ID,
				Description: tt.fields.Description,
				Completed:   tt.fields.Completed,
				CreatedAt:   tt.fields.CreatedAt,
			}

			todo.Update(tt.args.completed, tt.args.description)

			if todo.Completed != tt.args.completed {
				t1.Errorf("Update() = %v, want %v", todo.Completed, tt.args.completed)
			}
			if todo.Description != tt.args.description {
				t1.Errorf("Update() = %v, want %v", todo.Description, tt.args.description)
			}
		})
	}
}
