package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(userName string, password string) (todo.User, error)
}
type TodoList interface{}
type Todo interface{}

type Repository struct {
	Authorization Authorization
	TodoList      TodoList
	Todo          Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
