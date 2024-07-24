package repository

import (
	"github.com/jmoiron/sqlx"
	_ "managep"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
