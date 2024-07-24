package repository

import (
	"github.com/jmoiron/sqlx"
	_ "managep"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}
