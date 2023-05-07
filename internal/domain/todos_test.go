package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTodos_Add(t *testing.T) {
	type args struct {
		description string
	}
	tests := map[string]struct {
		l    Todos
		args args
		want *Todo
	}{
		"AddEmpty": {
			l:    Todos{},
			args: args{description: "test"},
			want: &Todo{
				ID:          uuid.New(),
				Description: "test",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
		},
		"AddNonEmpty": {
			l: Todos{
				{
					ID:          uuid.New(),
					Description: "test",
					Completed:   false,
					CreatedAt:   time.Now(),
				},
			},
			args: args{description: "test2"},
			want: &Todo{
				ID:          uuid.New(),
				Description: "test2",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			initialLength := len(tt.l)

			got := tt.l.Add(tt.args.description)

			if got.ID == uuid.Nil {
				t.Errorf("todo.ID = %v, want %v", got, tt.want)
			}
			if got.Description != tt.want.Description {
				t.Errorf("todo.Description = %v, want %v", got, tt.want)
			}
			if got.Completed != tt.want.Completed {
				t.Errorf("todo.Completed = %v, want %v", got, tt.want)
			}
			if got.CreatedAt.IsZero() {
				t.Errorf("todo.CreatedAt = %v, want %v", got, tt.want)
			}
			if len(tt.l) != initialLength+1 {
				t.Errorf("len(todos) = %v, want %v", got, tt.want)
			}
			if tt.l[initialLength] != got {
				t.Errorf("todos[%v] = %v, want %v", initialLength, tt.l[initialLength], tt.want)
			}
		})
	}
}

func TestTodos_All(t *testing.T) {
	var firstID = uuid.New()
	var first = &Todo{ID: firstID, Description: "first"}
	var secondID = uuid.New()
	var second = &Todo{ID: secondID, Description: "second"}
	var thirdID = uuid.New()
	var third = &Todo{ID: thirdID, Description: "third"}
	var fourthID = uuid.New()
	var fourth = &Todo{ID: fourthID, Description: "fourth"}
	tests := map[string]struct {
		l Todos

		want []*Todo
	}{
		"AllEmpty": {
			l:    Todos{},
			want: []*Todo{},
		},
		"AllNonEmpty": {
			l: Todos{
				first,
				second,
			},
			want: []*Todo{
				first,
				second,
			},
		},
		"AllMany": {
			l: Todos{
				first,
				second,
				third,
				fourth,
			},
			want: []*Todo{
				first,
				second,
				third,
				fourth,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.l.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_Get(t *testing.T) {
	var existingID = uuid.New()
	type args struct {
		id uuid.UUID
	}
	tests := map[string]struct {
		l    Todos
		args args
		want *Todo
	}{
		"GetEmpty": {
			l: Todos{},
			args: args{
				id: uuid.New(),
			},
			want: nil,
		},
		"GetNonEmpty": {
			l: Todos{
				{ID: uuid.New()},
			},
			args: args{
				id: uuid.New(),
			},
			want: nil,
		},
		"GetExisting": {
			l: Todos{
				{ID: existingID},
			},
			args: args{
				id: existingID,
			},
			want: &Todo{ID: existingID},
		},
		"GetExistingMultiple": {
			l: Todos{
				{ID: uuid.New()},
				{ID: existingID},
				{ID: uuid.New()},
			},
			args: args{
				id: existingID,
			},
			want: &Todo{ID: existingID},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.l.Get(tt.args.id)

			if tt.want == nil && got != nil {
				t.Errorf("todo = %v, want %v", got, tt.want)
			}
			// skip the rest of the test if we don't have a todo to compare
			if got == nil {
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("todo.ID = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_Remove(t *testing.T) {
	var existingID = uuid.New()
	type args struct {
		id uuid.UUID
	}
	tests := map[string]struct {
		l    Todos
		args args
	}{
		"RemoveEmpty": {
			l: Todos{},
			args: args{
				id: uuid.New(),
			},
		},
		"RemoveNonEmpty": {
			l: Todos{
				{ID: uuid.New()},
			},
			args: args{
				id: uuid.New(),
			},
		},
		"RemoveExisting": {
			l: Todos{
				{ID: existingID},
			},
			args: args{
				id: existingID,
			},
		},
		"RemoveExistingMultiple": {
			l: Todos{
				{ID: uuid.New()},
				{ID: existingID},
				{ID: uuid.New()},
			},
			args: args{
				id: existingID,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.l.Remove(tt.args.id)

			if tt.l.Get(tt.args.id) != nil {
				t.Errorf("todo = %v, want %v", tt.l.Get(tt.args.id), nil)
			}
		})
	}
}

func TestTodos_Reorder(t *testing.T) {
	var firstID = uuid.New()
	var first = &Todo{ID: firstID}
	var secondID = uuid.New()
	var second = &Todo{ID: secondID}
	var thirdID = uuid.New()
	var third = &Todo{ID: thirdID}
	var fourthID = uuid.New()
	var fourth = &Todo{ID: fourthID}

	type args struct {
		ids []uuid.UUID
	}
	tests := map[string]struct {
		l    Todos
		args args
		want []*Todo
	}{
		"ReorderEmpty": {
			l: Todos{},
			args: args{
				ids: []uuid.UUID{},
			},
			want: []*Todo{},
		},
		"ReorderNonEmpty": {
			l: Todos{
				first,
				second,
			},
			args: args{
				ids: []uuid.UUID{secondID, firstID},
			},
			want: []*Todo{
				second,
				first,
			},
		},
		"RecorderMany": {
			l: Todos{
				first,
				second,
				third,
				fourth,
			},
			args: args{
				ids: []uuid.UUID{fourthID, secondID, thirdID, firstID},
			},
			want: []*Todo{
				fourth,
				second,
				third,
				first,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.l.Reorder(tt.args.ids)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_Search(t *testing.T) {
	var firstID = uuid.New()
	var first = &Todo{ID: firstID, Description: "first"}
	var secondID = uuid.New()
	var second = &Todo{ID: secondID, Description: "second"}
	var thirdID = uuid.New()
	var third = &Todo{ID: thirdID, Description: "third"}
	var fourthID = uuid.New()
	var fourth = &Todo{ID: fourthID, Description: "fourth"}
	type args struct {
		search string
	}
	tests := map[string]struct {
		l    Todos
		args args
		want []*Todo
	}{
		"SearchEmpty": {
			l: Todos{},
			args: args{
				search: "first",
			},
			want: []*Todo{},
		},
		"SearchNonEmpty": {
			l: Todos{
				first,
				second,
			},
			args: args{
				search: "first",
			},
			want: []*Todo{
				first,
			},
		},
		"SearchMany": {
			l: Todos{
				first,
				second,
				third,
				fourth,
			},
			args: args{
				search: "ir",
			},
			want: []*Todo{
				first,
				third,
			},
		},
		"SearchManyNone": {
			l: Todos{
				first,
				second,
				third,
				fourth,
			},
			args: args{
				search: "z",
			},
			want: []*Todo{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.l.Search(tt.args.search); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodos_Update(t *testing.T) {
	var firstID = uuid.New()
	var secondID = uuid.New()
	var thirdID = uuid.New()
	type args struct {
		id          uuid.UUID
		completed   bool
		description string
	}
	tests := map[string]struct {
		l    Todos
		args args
		want *Todo
	}{
		"UpdateEmpty": {
			l: Todos{},
			args: args{
				id:          firstID,
				completed:   true,
				description: "first",
			},
			want: nil,
		},
		"UpdateNonEmpty": {
			l: Todos{
				{ID: firstID, Description: "first", Completed: false},
			},
			args: args{
				id:          secondID,
				completed:   true,
				description: "second",
			},
			want: nil,
		},
		"UpdateExisting": {
			l: Todos{
				{ID: firstID, Description: "first", Completed: false},
				{ID: secondID, Description: "second", Completed: false},
				{ID: thirdID, Description: "third", Completed: false},
			},
			args: args{
				id:          secondID,
				completed:   true,
				description: "SECOND",
			},
			want: &Todo{
				ID:          secondID,
				Description: "SECOND",
				Completed:   true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := tt.l.Update(tt.args.id, tt.args.completed, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
