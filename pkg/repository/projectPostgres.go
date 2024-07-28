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

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres(db *sqlx.DB) *ProjectPostgres {
	return &ProjectPostgres{db: db}
}

func (p *ProjectPostgres) IsOK(project *model.Project) bool {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE name=$1", projectTable)
	var exists bool
	err := p.db.QueryRow(query, project.Name).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true
		}
		log.Fatalf("Error checking project existence: %v", err)
	}
	return false
}

func (p *ProjectPostgres) IsProjectExists(id string) (bool, error) {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", projectTable)
	var exists bool
	err := p.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

func (p *ProjectPostgres) CreateProject(project *model.Project) (int, error) {
	if p.IsOK(project) == false {
		return http.StatusBadRequest, errors.New("project already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (name, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5)", projectTable)
	_, err := p.db.Exec(query, project.Name, project.Description, project.StartDate, project.FinishDate, project.Manager)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, nil
}

func (p *ProjectPostgres) GetProject() ([]model.Project, error) {
	var projects []model.Project
	var exists bool
	checkQuery := fmt.Sprintf("SELECT 1 FROM %s", projectTable)
	err := p.db.QueryRow(checkQuery).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]model.Project, 0), sql.ErrNoRows
		}
		return make([]model.Project, 0), sql.ErrNoRows
	}
	query := fmt.Sprintf("SELECT * FROM %s", projectTable)
	rows, err := p.db.Query(query)
	defer rows.Close()
	if err != nil {
		return make([]model.Project, 0), err
	}
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.StartDate, &project.FinishDate, &project.Manager); err != nil {
			return make([]model.Project, 0), err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (p *ProjectPostgres) GetProjectById(id string) (model.Project, error) {
	_, err := p.IsProjectExists(id)
	if err != nil {
		return model.Project{}, err
	}
	rows, err := p.db.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", projectTable), id)
	defer rows.Close()
	if err != nil {
		return model.Project{}, err
	}
	var project model.Project
	for rows.Next() {
		if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.FinishDate, &project.FinishDate, &project.Manager); err != nil {
			return model.Project{}, err
		}
	}
	return project, nil
}
func (p *ProjectPostgres) UpdateProject(project *model.Project, id string) (int, error) {
	_, err := p.IsProjectExists(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			return http.StatusNotFound, sql.ErrNoRows
		}
	}
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5 WHERE id = $6", projectTable)
	_, err = p.db.Exec(query, &project.Name, &project.Description, &project.StartDate, &project.FinishDate, &project.Manager, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (p *ProjectPostgres) DeleteProject(id string) (int, error) {
	_, err := p.IsProjectExists(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			return http.StatusNotFound, sql.ErrNoRows
		}
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", projectTable)
	_, err = p.db.Exec(query, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}
func (p *ProjectPostgres) SearchProject(query, queryType string) ([]model.Project, error) {
	var projects []model.Project
	var err error
	var q string
	var checkQuery string
	var exists bool
	if queryType == "title" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE name = $1", projectTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE name = $1", projectTable)
	} else if queryType == "manager" {
		q = fmt.Sprintf("SELECT * FROM %s WHERE manager_id = $1", projectTable)
		checkQuery = fmt.Sprintf("SELECT 1 FROM %s WHERE manager_id = $1", projectTable)
	}
	err = p.db.QueryRow(checkQuery, query).Scan(&exists)
	if err != nil {
		return make([]model.Project, 0), sql.ErrNoRows
	}
	rows, err := p.db.Query(q, query)
	if err != nil {
		return make([]model.Project, 0), err
	}
	defer rows.Close()
	for rows.Next() {
		var project model.Project
		if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.StartDate, &project.FinishDate, &project.Manager); err != nil {
			return make([]model.Project, 0), err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return make([]model.Project, 0), err
	}
	return projects, nil
}
func (p *ProjectPostgres) GetTasksForProject(id string) ([]model.Task, error) {
	var tasks []model.Task
	var exists bool
	_, err := p.IsProjectExists(id)
	if err != nil {
		errors.Is(err, sql.ErrNoRows)
		{
			return make([]model.Task, 1), sql.ErrNoRows
		}
	}
	checkQuery := fmt.Sprintf("SELECT 1 FROM %s WHERE project_id = $1", tasksTable)
	err = p.db.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]model.Task, 0), sql.ErrNoRows
		}
		return make([]model.Task, 0), err
	}
	rows, err := p.db.Query(fmt.Sprintf("SELECT * FROM %s WHERE project_id = $1", tasksTable), id)
	defer rows.Close()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]model.Task, 0), sql.ErrNoRows
		}
		return make([]model.Task, 0), err
	}
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
