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

func (p *ProjectPostgres) CreateProject(project *model.Project) (int, error) {
	if p.IsOK(project) == false {
		return 404, errors.New("project already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (name, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5)", projectTable)
	_, err := p.db.Exec(query, project.Name, project.Description, project.StartDate, project.FinishDate, project.Manager)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
