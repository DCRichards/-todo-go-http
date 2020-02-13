package rest

import (
	"errors"
	"github.com/dcrichards/todo-go-http/pkg/logger"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Server is an HTTP Server.
type Server struct {
	log         logger.Logger
	todoService todo.TodoService
	router      *httprouter.Router
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer creates a new Server instance.
func NewServer(options ...func(s *Server)) (*Server, error) {
	s := &Server{router: httprouter.New()}

	for _, o := range options {
		o(s)
	}

	if s.todoService == nil {
		return nil, errors.New("You must provide a TodoService")
	}

	if s.log == nil {
		var err error
		s.log, err = logger.New(
			logger.LogLevel(logger.LevelDebug),
			logger.LogOutput(logger.OutputStdout, nil),
			logger.LogFormat(logger.FormatText),
		)
		if err != nil {
			return nil, err
		}
	}

	s.routes()

	return s, nil
}

// TodoService sets the todo service for the server.
func TodoService(t todo.TodoService) func(*Server) {
	return func(s *Server) {
		s.todoService = t
	}
}

// Logger sets the logger for the server.
func Logger(l logger.Logger) func(*Server) {
	return func(s *Server) {
		s.log = l
	}
}
