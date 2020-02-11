package postgres

import (
	"errors"
	"fmt"
	"github.com/dcrichards/todo-go-http/pkg/todo"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type Todo struct {
	ID        int64  `pg:",pk"`
	Title     string `pg:",notnull"`
	Completed bool   `pg:",notnull"`
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

func (p *Postgres) GetAll() ([]todo.Todo, error) {
	rows := []Todo{}
	if err := p.db.Model(&rows).Select(); err != nil {
		return nil, err
	}

	// Convert to todo.Todo
	todos := []todo.Todo{}
	for _, r := range rows {
		todos = append(todos, todo.Todo{
			ID:        r.ID,
			Title:     r.Title,
			Completed: r.Completed,
		})
	}

	return todos, nil
}

func (p *Postgres) GetByID(id int64) (*todo.Todo, error) {
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
	return (*todo.Todo)(t), nil
}

func (p *Postgres) Create(todo *todo.Todo) (*todo.Todo, error) {
	return nil, nil
}

func (p *Postgres) Update(todo *todo.Todo) error {
	return nil
}

func (p *Postgres) Delete(id int64) error {
	return nil
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
