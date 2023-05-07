package todos

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/stackus/todos/internal/templates/pages"
	"github.com/stackus/todos/internal/templates/partials"
)

type (
	Handler interface {
		// Search : GET /todos
		Search(w http.ResponseWriter, r *http.Request)
		// Create : POST /todos
		Create(w http.ResponseWriter, r *http.Request)
		// Update : PATCH /todos/{todoId}
		// Update : POST /todos/{todoId}/edit
		Update(w http.ResponseWriter, r *http.Request)
		// Get : GET /todos/{todoId}
		Get(w http.ResponseWriter, r *http.Request)
		// Delete : DELETE /todos/{todoId}
		// Delete : POST /todos/{todoId}/delete
		Delete(w http.ResponseWriter, r *http.Request)
		// Sort : POST /todos/sort
		Sort(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Post("/", h.Create)
		r.Route("/{todoId}", func(r chi.Router) {
			r.Patch("/", h.Update)
			r.Post("/edit", h.Update)
			r.Get("/", h.Get)
			r.Delete("/", h.Delete)
			r.Post("/delete", h.Delete)
		})
		r.Post("/sort", h.Sort)
	})
}

func (h handler) Sort(w http.ResponseWriter, r *http.Request) {
	var todoIDs []uuid.UUID
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, id := range r.Form["id"] {
		var todoID uuid.UUID
		var err error
		if todoID, err = uuid.Parse(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		todoIDs = append(todoIDs, todoID)
	}
	if err := h.service.Sort(r.Context(), todoIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")
	todos, err := h.service.Search(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodos(todos).Render(r.Context(), w)
	default:
		err = pages.TodosPage(todos, search).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var description = r.Form.Get("description")

	todo, err := h.service.Add(r.Context(), description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodo(todo).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var completed = r.Form.Get("completed") == "true"
	var description = r.Form.Get("description")

	todo, err := h.service.Update(r.Context(), todoID, completed, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.RenderTodo(todo).Render(r.Context(), w)
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo, err := h.service.Get(r.Context(), todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		err = partials.EditTodoForm(todo).Render(r.Context(), w)
	default:
		err = pages.TodoPage(todo).Render(r.Context(), w)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "todoId")
	var todoID uuid.UUID
	var err error
	if todoID, err = uuid.Parse(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Remove(r.Context(), todoID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch isHTMX(r) {
	case true:
		_, err = w.Write([]byte(""))
	default:
		http.Redirect(w, r, "/", http.StatusFound)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isHTMX(r *http.Request) bool {
	// Check for "HX-Request" header
	if r.Header.Get("HX-Request") != "" {
		return true
	}

	return false
}
