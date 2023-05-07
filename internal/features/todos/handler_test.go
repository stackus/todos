package todos

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/stackus/todos/internal/domain"
	"github.com/stackus/todos/internal/templates/pages"
	"github.com/stackus/todos/internal/templates/partials"
)

func Test_handler_Create(t *testing.T) {
	var todo = &domain.Todo{
		ID:          uuid.New(),
		Description: "first",
		Completed:   false,
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
		"CreateHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("description=first"))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Add(context.Background(), "first").Return(todo, nil)
			},
			wantStatusCode: http.StatusFound,
			wantHeader: http.Header{
				"Location": []string{"/"},
			},
			wantView: nil,
		},
		"CreateHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("description=first"))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					req.Header.Set("HX-Request", "true")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Add(context.Background(), "first").Return(todo, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/html; charset=utf-8"},
			},
			wantView: partials.RenderTodo(todo),
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

			h.Create(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Create() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Create() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Create() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Create() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}

func Test_handler_Delete(t *testing.T) {
	var todoID = uuid.New()
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
		"DeleteHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/", nil)
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Remove(mock.AnythingOfType("*context.valueCtx"), todoID).Return(nil)
			},
			wantStatusCode: http.StatusFound,
			wantHeader: http.Header{
				"Location": []string{"/"},
			},
			wantView: nil,
		},
		"DeleteHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodDelete, "/", nil)
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
					req.Header.Set("HX-Request", "true")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Remove(mock.AnythingOfType("*context.valueCtx"), todoID).Return(nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/plain; charset=utf-8"},
			},
			wantView: nil,
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

			h.Delete(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Delete() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Delete() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Delete() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Delete() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}

func Test_handler_Get(t *testing.T) {
	var todoID = uuid.New()
	var todo = &domain.Todo{
		ID:          todoID,
		Description: "test",
		Completed:   false,
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
		"GetHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Get(mock.AnythingOfType("*context.valueCtx"), todoID).Return(todo, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/html; charset=utf-8"},
			},
			wantView: pages.TodoPage(todo),
		},
		"GetHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
					req.Header.Set("HX-Request", "true")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Get(mock.AnythingOfType("*context.valueCtx"), todoID).Return(todo, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/html; charset=utf-8"},
			},
			wantView: partials.EditTodoForm(todo),
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

			h.Get(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Get() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Get() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Get() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Get() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}

func Test_handler_Search(t *testing.T) {
	var todo = &domain.Todo{
		ID:          uuid.New(),
		Description: "test",
		Completed:   false,
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
		"SearchHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/?search=test", nil)
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Search(mock.AnythingOfType("*context.emptyCtx"), "test").Return([]*domain.Todo{todo}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/html; charset=utf-8"},
			},
			wantView: pages.TodosPage([]*domain.Todo{todo}, "test"),
		},
		"SearchHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodGet, "/?search=test", nil)
					req.Header.Set("HX-Request", "true")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Search(mock.AnythingOfType("*context.emptyCtx"), "test").Return([]*domain.Todo{todo}, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/plain; charset=utf-8"},
			},
			wantView: partials.RenderTodos([]*domain.Todo{todo}),
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

			h.Search(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Search() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Search() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Search() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Search() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}

func Test_handler_Sort(t *testing.T) {
	var todoIDs = []uuid.UUID{uuid.New(), uuid.New()}
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
		"SortHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					ids := make([]string, len(todoIDs))
					for i, id := range todoIDs {
						ids[i] = id.String()
					}
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(url.Values{"id": ids}.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Sort(mock.AnythingOfType("*context.emptyCtx"), todoIDs).Return(nil)
			},
			wantStatusCode: http.StatusFound,
			wantHeader: http.Header{
				"Location": []string{"/"},
			},
			wantView: nil,
		},
		"SortHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					ids := make([]string, len(todoIDs))
					for i, id := range todoIDs {
						ids[i] = id.String()
					}
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(url.Values{"id": ids}.Encode()))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					req.Header.Set("HX-Request", "true")
					return req
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Sort(mock.AnythingOfType("*context.emptyCtx"), todoIDs).Return(nil)
			},
			wantStatusCode: http.StatusNoContent,
			wantHeader:     http.Header{},
			wantView:       nil,
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

			h.Sort(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Sort() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Sort() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Sort() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Sort() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}

func Test_handler_Update(t *testing.T) {
	var todoID = uuid.New()
	var updated = &domain.Todo{
		ID:          todoID,
		Description: "first",
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
		"UpdateHTML": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("description=first&completed=true"))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Update(mock.AnythingOfType("*context.valueCtx"), todoID, true, "first").Return(updated, nil)
			},
			wantStatusCode: http.StatusFound,
			wantHeader: http.Header{
				"Location": []string{"/"},
			},
			wantView: nil,
		},
		"UpdateHTMX": {
			args: args{
				w: httptest.NewRecorder(),
				r: func() *http.Request {
					req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("description=first&completed=true"))
					req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
					req.Header.Set("HX-Request", "true")
					rCtx := chi.NewRouteContext()
					rCtx.URLParams.Add("todoId", todoID.String())
					return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))
				}(),
			},
			mock: func(f fields) {
				f.service.EXPECT().Update(mock.AnythingOfType("*context.valueCtx"), todoID, true, "first").Return(updated, nil)
			},
			wantStatusCode: http.StatusOK,
			wantHeader: http.Header{
				"Content-Type": []string{"text/html; charset=utf-8"},
			},
			wantView: partials.RenderTodo(updated),
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

			h.Update(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("handler.Update() StatusCode = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}
			if !reflect.DeepEqual(res.Header, tt.wantHeader) {
				t.Errorf("handler.Update() Header = %v, want %v", res.Header, tt.wantHeader)
			}
			gotBuffer := new(bytes.Buffer)
			_, _ = gotBuffer.ReadFrom(res.Body)
			if tt.wantView == nil {
				if strings.TrimSpace(gotBuffer.String()) != "" {
					t.Errorf("handler.Update() Body = %v, want %v", gotBuffer.String(), "")
				}
				return
			}
			wantBuffer := new(bytes.Buffer)
			_ = tt.wantView.Render(context.Background(), wantBuffer)
			if gotBuffer.String() != wantBuffer.String() {
				t.Errorf("handler.Update() Body = %v, want %v", gotBuffer.String(), wantBuffer.String())
			}
		})
	}
}
