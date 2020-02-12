package testutils

import (
	todoService "github.com/dcrichards/todo-go-http/pkg/todo"
)

var Todos = []todoService.Todo{
	{ID: 0, Title: "Clean my bath", Completed: false},
	{ID: 1, Title: "Return my dog", Completed: true},
	{ID: 2, Title: "Straighen out my bird", Completed: true},
	{ID: 3, Title: "Sell me cigarettes", Completed: false},
}

type MockRepository struct{}

func (m *MockRepository) GetAll() ([]todoService.Todo, error) {
	return Todos, nil
}

func (m *MockRepository) GetByID(id int64) (*todoService.Todo, error) {
	if id < 0 || int(id) > len(Todos)-1 {
		return nil, nil
	}

	return &Todos[id], nil
}

func (m *MockRepository) Create(todo *todoService.Todo) (*todoService.Todo, error) {
	todo.ID = 42
	return todo, nil
}

func (m *MockRepository) Update(todo *todoService.Todo) error {
	return nil
}

func (m *MockRepository) Delete(id int64) error {
	return nil
}
