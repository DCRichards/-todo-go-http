package testutils

import (
	"errors"
	todoService "github.com/dcrichards/todo-go-http/pkg/todo"
)

// GoodTodoService is a well-behaved todo service which behaves normally.
type GoodTodoService struct{}

// BadTodoService is a badly behaved service which returns errors.
type BadTodoService struct{}

func (m *GoodTodoService) GetAll() ([]todoService.Todo, error) {
	return Todos, nil
}

func (m *GoodTodoService) GetByID(id int64) (*todoService.Todo, error) {
	if id < 0 || int(id) > len(Todos)-1 {
		return nil, nil
	}

	return &Todos[id], nil
}

func (m *GoodTodoService) Create(todo *todoService.Todo) (*todoService.Todo, error) {
	todo.ID = 12
	return todo, nil
}

func (m *GoodTodoService) Update(todo *todoService.Todo) error {
	return nil
}

func (m *GoodTodoService) Delete(id int64) error {
	return nil
}

func (m *BadTodoService) GetAll() ([]todoService.Todo, error) {
	return nil, errors.New("Get it yourself, hotshot!")
}

func (m *BadTodoService) GetByID(id int64) (*todoService.Todo, error) {
	return nil, errors.New("Can't be bothered to find ID")
}

func (m *BadTodoService) Create(todo *todoService.Todo) (*todoService.Todo, error) {
	return nil, errors.New("Um probably no room or something")
}

func (m *BadTodoService) Update(todo *todoService.Todo) error {
	return errors.New("Broken.")
}

func (m *BadTodoService) Delete(id int64) error {
	return errors.New("Can't delete. It's there forever. Facebook style.")
}
