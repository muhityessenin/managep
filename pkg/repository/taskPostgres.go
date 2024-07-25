package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	_ "managep"
	"managep/pkg/model"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) Check(task *model.Task) bool {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE name=$1", tasksTable)
	var exists bool
	err := t.db.QueryRow(query, task.Name).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true
		}
		log.Fatalf("Error checking task existence: %v", err)
	}
	return false
}

func (t *TaskPostgres) CreateTask(task *model.Task) (int, error) {
	if t.Check(task) == true {
		return 404, errors.New("user is already registered")
	}
	query := fmt.Sprintf("INSERT INTO %s (name, description, priority, state, responsible_person_id, project_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", tasksTable)
	_, err := t.db.Exec(query, task.Name, task.Description, task.Priority, task.State, task.ResponsiblePerson, task.Project, task.CreatedAt, task.FinishedAt)
	if err != nil {
		return 404, err
	}
	return 201, nil
}
