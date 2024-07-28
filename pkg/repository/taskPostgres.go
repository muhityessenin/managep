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

func (t *TaskPostgres) IsTaskExist(id string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", tasksTable)
	var exists bool
	err := t.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, sql.ErrNoRows
		}
		return false, err
	}
	return true, nil
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
	return http.StatusCreated, nil
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
	if len(tasks) == 0 {
		return nil, sql.ErrNoRows
	}
	return tasks, nil
}

func (t *TaskPostgres) GetTaskById(id string) (model.Task, error) {
	_, err := t.IsTaskExist(id)
	if err != nil {
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
	_, err := t.IsTaskExist(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			return http.StatusNotFound, err
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
	_, err := t.IsTaskExist(id)
	if err != nil {
		return http.StatusNotFound, err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tasksTable)
	_, err = t.db.Exec(query, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (t *TaskPostgres) SearchTask(query, queryType string) ([]model.Task, error) {
	var tasks []model.Task
	var err error
	var q string
	var checkQuery string
	var exists bool
	if queryType == "title" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE name = $1", tasksTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE name = $1", tasksTable)
	} else if queryType == "status" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE state = $1", tasksTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE state = $1", tasksTable)
	} else if queryType == "priority" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE priority = $1", tasksTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE priority = $1", tasksTable)
	} else if queryType == "assignee" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE responsible_person_id = $1", tasksTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE responsible_person_id = $1", tasksTable)
	} else if queryType == "project" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE project_id = $1", tasksTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE project_id = $1", tasksTable)
	}
	err = t.db.QueryRow(checkQuery, query).Scan(&exists)
	if err != nil {
		return make([]model.Task, 0), sql.ErrNoRows
	}
	rows, err := t.db.Query(q, query)
	if err != nil {
		return make([]model.Task, 0), err
	}
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Priority, &task.State, &task.ResponsiblePerson, &task.Project, &task.CreatedAt, &task.FinishedAt); err != nil {
			return make([]model.Task, 0), err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return make([]model.Task, 0), err
	}
	return tasks, nil
}
