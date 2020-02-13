package rest

import (
	"encoding/json"
	"fmt"
	"github.com/dcrichards/todo-go-http/pkg/logger"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoResponse struct {
	Todo *todo.Todo `json:"todo,omitempty"`
}

type TodosResponse struct {
	Todos []todo.Todo `json:"todos"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s *Server) routes() {
	s.router.GET("/todo", s.getTodos())
	s.router.GET("/todo/:id", s.getTodoByID())
	s.router.POST("/todo", s.createTodo())
	s.router.PUT("/todo/:id", s.updateTodo())
	s.router.DELETE("/todo/:id", s.deleteTodo())
}

func (s *Server) getTodos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		todos, err := s.todoService.GetAll()
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to get todos.", err)
			return
		}

		response := &TodosResponse{todos}
		s.handleJSON(w, response)
	}
}

func (s *Server) getTodoByID() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		q := p.ByName("id")
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			s.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %s.", q), nil)
			return
		}

		todo, err := s.todoService.GetByID(id)
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to get todo.", err)
			return
		}

		if todo == nil {
			s.handleError(w, http.StatusNotFound, fmt.Sprintf("No todos found with ID: %s.", q), nil)
			return
		}

		s.handleJSON(w, &TodoResponse{todo})
	}
}

func (s *Server) createTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var body TodoRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to process request.", err)
			return
		}

		if body.Title == "" {
			s.handleError(w, http.StatusBadRequest, "You must specify a valid title", nil)
			return
		}

		created, err := s.todoService.Create(&todo.Todo{
			Title:     body.Title,
			Completed: body.Completed,
		})
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to create todo.", err)
			return
		}

		s.handleJSON(w, &TodoResponse{created})
	}
}

func (s *Server) updateTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		q := p.ByName("id")
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			s.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %s.", q), nil)
			return
		}

		var body TodoRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to process request.", err)
			return
		}

		err = s.todoService.Update(&todo.Todo{
			ID:        id,
			Title:     body.Title,
			Completed: body.Completed,
		})
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to update todo.", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) deleteTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		q := p.ByName("id")
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			s.handleError(w, http.StatusBadRequest, fmt.Sprintf("Invalid ID: %s.", q), nil)
			return
		}

		if err := s.todoService.Delete(id); err != nil {
			s.handleError(w, http.StatusBadRequest, "Unable to delete todo.", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *Server) handleError(w http.ResponseWriter, status int, msg string, err error) {
	if err != nil {
		s.log.Error(err, logger.Meta{
			"statusCode": status,
		})
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&ErrorResponse{msg}); err != nil {
		s.log.Error(err)
	}
}

func (s *Server) handleJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		s.handleError(w, http.StatusInternalServerError, "Unable to process response.", err)
	}
}
