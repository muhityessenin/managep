package service

import (
	todo "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(userName string, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface{}
type Todo interface{}

type Service struct {
	Authorization Authorization
	TodoList      TodoList
	Todo          Todo
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
