package todo_test

import (
	"github.com/dcrichards/todo-go-http/internal/testutils"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"reflect"
	"testing"
)

func TestGetAll(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})

	todos, err := s.GetAll()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testutils.Todos, todos) {
		t.Errorf("Expected:\n%+v\nBut got:\n%+v\n", testutils.Todos, todos)
	}
}

func TestGetByID_Valid(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})

	testCases := []struct {
		ID       int64
		Expected *todo.Todo
	}{
		{ID: 1, Expected: &testutils.Todos[1]},
		{ID: 2, Expected: &testutils.Todos[2]},
		{ID: 3, Expected: &testutils.Todos[3]},
	}

	for _, c := range testCases {
		actual, err := s.GetByID(c.ID)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(c.Expected, actual) {
			t.Errorf("Expected:\n%+v\nBut got:\n%+v\n", c.Expected, actual)
		}
	}
}

func TestGetByID_Invalid(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})

	testCases := []int64{-1, 5}
	for _, id := range testCases {
		actual, err := s.GetByID(id)
		if err != nil {
			t.Error(err)
		}
		if actual != nil {
			t.Errorf("Expected GetByID(%d) to return nil", id)
		}
	}
}

func TestCreate(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})
	create := &todo.Todo{Title: "Simmer down", Completed: true}

	actual, err := s.Create(create)
	if err != nil {
		t.Error(err)
	}

	if actual.ID == 0 {
		t.Error("Expected todo to be returned with assigned ID")
	}

	if actual.Title != create.Title {
		t.Errorf("Expected Title '%s', but got '%s'", create.Title, actual.Title)
	}

	if actual.Title != create.Title {
		t.Errorf("Expected Completed '%t', but got '%t'", create.Completed, actual.Completed)
	}
}

func TestUpdate(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})
	if err := s.Update(&testutils.Todos[1]); err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	s := todo.NewService(&testutils.MockRepository{})
	if err := s.Delete(10); err != nil {
		t.Error(err)
	}
}
