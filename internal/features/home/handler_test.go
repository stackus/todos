package home

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"

	"github.com/stackus/todos/internal/domain"
	"github.com/stackus/todos/internal/templates/pages"
)

func Test_handler_Home(t *testing.T) {
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
		service *MockService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := map[string]struct {
		args           args
		mock           func(f fields)
		wantStatusCode int
		wantHeader     http.Header
		wantView       templ.Component
	}{
		"EmptyList": {
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			mock: func(f fields) {
				f.service.EXPECT().List(context.Background()).Return([]*domain.Todo{}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader:     http.Header{"Content-Type": []string{"text/html; charset=utf-8"}},
			wantView:       pages.HomePage([]*domain.Todo{}),
		},
		"NonEmptyList": {
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/", nil),
			},
			mock: func(f fields) {
				f.service.EXPECT().List(context.Background()).Return([]*domain.Todo{firstTodo, secondTodo, thirdTodo}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader:     http.Header{"Content-Type": []string{"text/html; charset=utf-8"}},
			wantView:       pages.HomePage([]*domain.Todo{firstTodo, secondTodo, thirdTodo}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			f := fields{
				service: NewMockService(t),
			}
			h := handler{
				service: f.service,
			}
			if tt.mock != nil {
				tt.mock(f)
			}

			h.Home(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Home() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Home() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Home() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}
