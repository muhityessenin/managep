package repository

import (
	"github.com/jmoiron/sqlx"
	_ "managep"
)

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}
