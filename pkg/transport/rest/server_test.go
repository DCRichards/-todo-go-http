package rest_test

import (
	"encoding/json"
	"fmt"
	"github.com/dcrichards/todo-go-http/internal/testutils"
	"github.com/dcrichards/todo-go-http/pkg/transport/rest"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func newRequest(s *rest.Server, url string) (*http.Response, error) {
	server := httptest.NewServer(s)
	defer server.Close()

	return server.Client().Get(fmt.Sprintf("%s/%s", server.URL, url))
}

func TestGetTodos_200(t *testing.T) {
	s, err := rest.NewServer(rest.TodoService(&testutils.GoodTodoService{}))
	if err != nil {
		t.Error(err)
	}

	res, err := newRequest(s, "todo")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 but got %d", res.StatusCode)
	}

	expected := rest.TodosResponse{
		Todos: testutils.Todos,
	}

	var actual rest.TodosResponse
	if err := json.NewDecoder(res.Body).Decode(&actual); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected:\n%+v\nBut got:\n%+v\n", expected, actual)
	}
}

func TestGetTodos_500(t *testing.T) {
	s, err := rest.NewServer(rest.TodoService(&testutils.BadTodoService{}))
	if err != nil {
		t.Error(err)
	}

	server := httptest.NewServer(s)
	defer server.Close()

	res, err := newRequest(s, "todo")
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500 but got %d", res.StatusCode)
	}

	var actual rest.ErrorResponse
	if err := json.NewDecoder(res.Body).Decode(&actual); err != nil {
		t.Error(err)
	}

	if actual.Error == "" {
		t.Error("Expected error response to contain error message")
	}
}
