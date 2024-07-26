package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	_ "managep"
	"managep/pkg/model"
	"net/http"
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
	if t.Check(task) == false {
		return http.StatusBadRequest, errors.New("task is already registered")
	}
	query := fmt.Sprintf("INSERT INTO %s (name, description, priority, state, responsible_person_id, project_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", tasksTable)
	_, err := t.db.Exec(query, task.Name, task.Description, task.Priority, task.State, task.ResponsiblePerson, task.Project, task.CreatedAt, task.FinishedAt)
	if err != nil {
		fmt.Printf("Error creating task: %v", err.Error())
		return http.StatusBadRequest, err
	}
	return 201, nil
}

func (t *TaskPostgres) GetTask() ([]model.Task, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tasksTable)
	rows, err := t.db.Query(query)
	if err != nil {
		return []model.Task{}, err
	}
	defer rows.Close()
	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.State, &task.ResponsiblePerson, &task.Project, &task.CreatedAt, &task.FinishedAt); err != nil {
			return []model.Task{}, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
}

func (t *TaskPostgres) GetTaskById(id string) (model.Task, error) {
	checkQuery := fmt.Sprintf("SELECT id FROM %s WHERE id = $1", tasksTable)
	err := t.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, errors.New("task not found")
		}
		return model.Task{}, err
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tasksTable)
	rows, err := t.db.Query(query, id)
	defer rows.Close()
	var task model.Task
	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.State, &task.ResponsiblePerson, &task.Project, &task.CreatedAt, &task.FinishedAt); err != nil {
			return model.Task{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t *TaskPostgres) UpdateTask(task *model.Task, id string) (int, error) {
	checkQuery := fmt.Sprintf("SELECT id FROM %s WHERE id = $1", tasksTable)
	err := t.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 404, errors.New("task not found")
		}
		return http.StatusBadRequest, err
	}
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, priority = $3, state = $4, responsible_person_id = $5, project_id = $6, start_date = $7, end_date = $8 WHERE id = $9", tasksTable)
	_, err = t.db.Exec(query, task.Name, task.Description, task.Priority, task.State, task.ResponsiblePerson, task.Project, task.CreatedAt, task.FinishedAt, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (t *TaskPostgres) DeleteTask(id string) (int, error) {
	checkQuery := fmt.Sprintf("SELECT id FROM %s WHERE id = $1", tasksTable)
	err := t.db.QueryRow(checkQuery, id).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 404, errors.New("task not found")
		}
		return 404, err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tasksTable)
	_, err = t.db.Exec(query, id)
	if err != nil {
		return 404, err
	}
	return http.StatusOK, nil
}
