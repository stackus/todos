package todos

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/stackus/todos/internal/domain"
)

func Test_service_Add(t *testing.T) {
	var todo = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx         context.Context
		description string
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		want    *domain.Todo
		wantErr bool
	}{
		"Add": {
			args: args{
				ctx:         context.Background(),
				description: "first",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Add("first").Return(todo)
			},
			want:    todo,
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
			got, err := s.Add(tt.args.ctx, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	var todo = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		want    *domain.Todo
		wantErr bool
	}{
		"GetNil": {
			args: args{
				ctx: context.Background(),
				id:  todo.ID,
			},
			mock: func(f fields) {
				f.todos.EXPECT().Get(todo.ID).Return(nil)
			},
			want:    nil,
			wantErr: false,
		},
		"GetMatch": {
			args: args{
				ctx: context.Background(),
				id:  todo.ID,
			},
			mock: func(f fields) {
				f.todos.EXPECT().Get(todo.ID).Return(todo)
			},
			want:    todo,
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
			got, err := s.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Remove(t *testing.T) {
	var todoID = uuid.New()
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		wantErr bool
	}{
		"Remove": {
			args: args{
				ctx: context.Background(),
				id:  todoID,
			},
			mock: func(f fields) {
				f.todos.EXPECT().Remove(todoID)
			},
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
			if err := s.Remove(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Search(t *testing.T) {
	var first = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	var third = &domain.Todo{
		ID:          uuid.New(),
		Description: "third",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx    context.Context
		search string
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		want    []*domain.Todo
		wantErr bool
	}{
		"SearchEmpty": {
			args: args{
				ctx:    context.Background(),
				search: "fourth",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Search("fourth").Return([]*domain.Todo{})
			},
			want:    []*domain.Todo{},
			wantErr: false,
		},
		"SearchOne": {
			args: args{
				ctx:    context.Background(),
				search: "first",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Search("first").Return([]*domain.Todo{first})
			},
			want:    []*domain.Todo{first},
			wantErr: false,
		},
		"SearchMany": {
			args: args{
				ctx:    context.Background(),
				search: "ir",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Search("ir").Return([]*domain.Todo{first, third})
			},
			want:    []*domain.Todo{first, third},
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
			got, err := s.Search(tt.args.ctx, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Sort(t *testing.T) {
	var first = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	var second = &domain.Todo{
		ID:          uuid.New(),
		Description: "second",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	var third = &domain.Todo{
		ID:          uuid.New(),
		Description: "third",
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx context.Context
		ids []uuid.UUID
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		wantErr bool
	}{
		"SortEmpty": {
			args: args{
				ctx: context.Background(),
				ids: []uuid.UUID{},
			},
			mock: func(f fields) {
				f.todos.EXPECT().Reorder([]uuid.UUID{}).Return([]*domain.Todo{})
			},
			wantErr: false,
		},
		"SortMany": {
			args: args{
				ctx: context.Background(),
				ids: []uuid.UUID{third.ID, second.ID, first.ID},
			},
			mock: func(f fields) {
				f.todos.EXPECT().Reorder([]uuid.UUID{third.ID, second.ID, first.ID}).Return([]*domain.Todo{third, second, first})
			},
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
			if err := s.Sort(tt.args.ctx, tt.args.ids); (err != nil) != tt.wantErr {
				t.Errorf("Sort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Update(t *testing.T) {
	var todoID = uuid.New()
	var updated = &domain.Todo{
		ID:          todoID,
		Description: "updated",
		Completed:   true,
		CreatedAt:   time.Now(),
	}
	type fields struct {
		todos *domain.MockTodoRepository
	}
	type args struct {
		ctx         context.Context
		id          uuid.UUID
		completed   bool
		description string
	}
	tests := map[string]struct {
		args    args
		mock    func(f fields)
		want    *domain.Todo
		wantErr bool
	}{
		"UpdateNotFound": {
			args: args{
				ctx:         context.Background(),
				id:          todoID,
				completed:   true,
				description: "updated",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Update(todoID, true, "updated").Return(nil)
			},
			want:    nil,
			wantErr: false,
		},
		"UpdateFound": {
			args: args{
				ctx:         context.Background(),
				id:          todoID,
				completed:   true,
				description: "updated",
			},
			mock: func(f fields) {
				f.todos.EXPECT().Update(todoID, true, "updated").Return(updated)
			},
			want:    updated,
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
			got, err := s.Update(tt.args.ctx, tt.args.id, tt.args.completed, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
