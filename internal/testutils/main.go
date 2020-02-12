package testutils

import (
	todoService "github.com/dcrichards/todo-go-http/pkg/todo"
)

// Todos is a collection of example Todo items for testing.
var Todos = []todoService.Todo{
	{ID: 0, Title: "Clean my bath", Completed: false},
	{ID: 1, Title: "Return my dog", Completed: true},
	{ID: 2, Title: "Straighen out my bird", Completed: true},
	{ID: 3, Title: "Sell me cigarettes", Completed: false},
}
