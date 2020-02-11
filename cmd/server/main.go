package main

import (
	"fmt"
	"github.com/dcrichards/todo-go-http/internal/environment"
	"github.com/dcrichards/todo-go-http/pkg/persistence/postgres"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/dcrichards/todo-go-http/pkg/transport/rest"
	"log"
	"net/http"
)

func main() {
	env, err := environment.Get()
	if err != nil {
		panic(err)
	}

	db, err := postgres.NewPostgres(&postgres.ConnectionParams{
		Host:     env.PostgresHost,
		Port:     env.PostgresPort,
		Username: env.PostgresUser,
		Password: env.PostgresPassword,
		Database: env.PostgresDb,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	todoService := todo.NewService(db)
	server, err := rest.NewServer(rest.TodoService(todoService))
	if err != nil {
		panic(err)
	}

	log.Printf("🌎 Server running on :%d\n", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), server))
}
