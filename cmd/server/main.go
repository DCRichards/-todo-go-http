package main

import (
	"fmt"
	"github.com/dcrichards/todo-go-http/internal/environment"
	"github.com/dcrichards/todo-go-http/pkg/logger"
	"github.com/dcrichards/todo-go-http/pkg/persistence/postgres"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/dcrichards/todo-go-http/pkg/transport/rest"
	"net/http"
	"os"
	"strings"
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

	log, err := getLogger(env.LogLevel, env.LogFormat, env.LogFile)
	if err != nil {
		panic(err)
	}

	todoService := todo.NewService(db)
	server, err := rest.NewServer(
		rest.TodoService(todoService),
		rest.Logger(log),
	)
	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("🌎 Server running on :%d\n", env.Port))
	log.Error(http.ListenAndServe(fmt.Sprintf(":%d", env.Port), server))
}

func getLogger(level string, format string, logFile string) (*logger.Log, error) {
	logLevel := logger.LogLevel(logger.LevelInfo)
	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel = logger.LogLevel(logger.LevelDebug)
	case "ERROR":
		logLevel = logger.LogLevel(logger.LevelError)
	}

	logFmt := logger.LogFormat(logger.FormatText)
	switch strings.ToUpper(format) {
	case "JSON":
		logFmt = logger.LogFormat(logger.FormatJSON)
	}

	logOutput := logger.LogOutput(logger.OutputStdout, nil)
	if logFile != "" {
		logFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		logOutput = logger.LogOutput(logger.OutputFile, logFile)
	}

	return logger.New(logLevel, logFmt, logOutput)
}
