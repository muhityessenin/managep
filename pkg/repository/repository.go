package repository

import (
	"github.com/jmoiron/sqlx"
	_ "managep"
)

type Users interface {
}
type Tasks interface {
}
type Projects interface {
}

type Repository struct {
	UserService    Users
	TaskService    Tasks
	ProjectService Projects
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserService:    NewUserPostgres(db),
		TaskService:    NewTaskPostgres(db),
		ProjectService: NewProjectPostgres(db),
	}
}
