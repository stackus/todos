package home

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stackus/todos/internal/domain"
)

func Test_service_List(t *testing.T) {
	var firstTodo = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	var secondTodo = &domain.Todo{
		ID:          uuid.New(),
		Description: "second",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	var thirdTodo = &domain.Todo{
		ID:          uuid.New(),
		Description: "third",
		Completed:   true,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		want    []*domain.Todo
		wantErr bool
	}{
		"ListEmpty": {
			args: args{
				ctx: context.Background(),
			},
			mock: func(f fields) {
				f.todos.EXPECT().All().Return([]*domain.Todo{})
			},
			want:    []*domain.Todo{},
			wantErr: false,
		},
		"ListNonEmpty": {
			args: args{
				ctx: context.Background(),
			},
			mock: func(f fields) {
				f.todos.EXPECT().All().Return([]*domain.Todo{firstTodo, secondTodo, thirdTodo})
			},
			want:    []*domain.Todo{firstTodo, secondTodo, thirdTodo},
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f := fields{
				todos: domain.NewMockTodoRepository(t),
			}
			s := service{
				todos: f.todos,
			}
			if tt.mock != nil {
				tt.mock(f)
			}
			got, err := s.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
