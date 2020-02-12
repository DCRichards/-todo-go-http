package postgres

import (
	"errors"
	"fmt"
	todoService "github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type Todo struct {
	ID        int64  `pg:",pk"`
	Title     string `pg:",notnull"`
	Completed bool   `pg:",notnull,use_zero"`
}

type Postgres struct {
	db *pg.DB
}

type ConnectionParams struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) GetAll() ([]todoService.Todo, error) {
	rows := []Todo{}
	if err := p.db.Model(&rows).Select(); err != nil {
		return nil, err
	}

	// Convert to todoService.Todo
	todos := []todoService.Todo{}
	for _, r := range rows {
		todos = append(todos, todoService.Todo{
			ID:        r.ID,
			Title:     r.Title,
			Completed: r.Completed,
		})
	}

	return todos, nil
}

func (p *Postgres) GetByID(id int64) (*todoService.Todo, error) {
	t := &Todo{ID: id}
	err := p.db.Select(t)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	// Not found.
	if t.Title == "" {
		return nil, nil
	}

	// We can cast here as the fields are the same.
	return (*todoService.Todo)(t), nil
}

func (p *Postgres) Create(todo *todoService.Todo) (*todoService.Todo, error) {
	t := &Todo{
		Title:     todo.Title,
		Completed: todo.Completed,
	}

	_, err := p.db.Model(t).Insert()
	if err != nil {
		return nil, err
	}

	return (*todoService.Todo)(t), nil
}

func (p *Postgres) Update(todo *todoService.Todo) error {
	t := &Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}

	return p.db.Update(t)
}

func (p *Postgres) Delete(id int64) error {
	t := &Todo{ID: id}
	return p.db.Delete(t)
}

func NewPostgres(params *ConnectionParams) (*Postgres, error) {
	if params == nil {
		return nil, errors.New("db: You must specify connection params")
	}

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", params.Host, params.Port),
		User:     params.Username,
		Password: params.Password,
		Database: params.Database,
	})

	if err := db.CreateTable(&Todo{}, &orm.CreateTableOptions{IfNotExists: true}); err != nil {
		return nil, err
	}

	return &Postgres{db}, nil
}
