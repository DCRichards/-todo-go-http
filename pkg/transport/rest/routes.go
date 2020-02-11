package rest

import (
	"encoding/json"
	"fmt"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (s *Server) routes() {
	s.router.GET("/todo", s.getTodos())
	s.router.GET("/todo/:id", s.getTodoByID())
	s.router.POST("/todo", s.createTodo())
	s.router.PUT("/todo/:id", s.updateTodo())
	s.router.DELETE("/todo/:id", s.deleteTodo())
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TodoResponse struct {
	Todo *todo.Todo `json:"todo,omitempty"`
}

type TodosResponse struct {
	Todos []todo.Todo `json:"todos"`
}

func (s *Server) getTodos() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		todos, err := s.todoService.GetAll()
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to get todos.")
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
			s.handleError(w, http.StatusInternalServerError, fmt.Sprintf("Invalid ID: %s", q))
			return
		}

		todo, err := s.todoService.GetByID(id)
		if err != nil {
			s.handleError(w, http.StatusInternalServerError, "Unable to get todo.")
			return
		}

		if todo == nil {
			s.handleError(w, http.StatusNotFound, fmt.Sprintf("No todos found with ID: %s", q))
			return
		}

		response := &TodoResponse{todo}
		s.handleJSON(w, response)
	}
}

func (s *Server) createTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// TODO s.todoService.stuff()
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (s *Server) updateTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// TODO s.todoService.stuff()
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (s *Server) deleteTodo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// TODO s.todoService.stuff()
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (s *Server) handleError(w http.ResponseWriter, status int, message string) {
	// TODO: Log the error.
	w.WriteHeader(status)
	// Deliberately ignore error here as if that fails,
	// we've got no other to inform the client anyway.
	json.NewEncoder(w).Encode(&ErrorResponse{message})
}

func (s *Server) handleJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		s.handleError(w, http.StatusInternalServerError, "Unable to process response.")
	}
}
