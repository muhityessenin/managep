package repository

import (
	"github.com/jmoiron/sqlx"
	_ "managep"
	"managep/pkg/model"
)

type Users interface {
	GetUser() ([]model.User, error)
	CreateUser(user *model.User) (int, error)
	GetUserById(id string) (model.User, error)
	UpdateUser(user *model.User, id string) (int, error)
	DeleteUser(id string) (int, error)
	GetTasksForUser(id string) ([]model.Task, error)
	SearchUser(query, queryType string) (model.User, error)
}
type Tasks interface {
	GetTask() ([]model.Task, error)
	CreateTask(task *model.Task) (int, error)
	GetTaskById(id string) (model.Task, error)
	UpdateTask(task *model.Task, id string) (int, error)
	DeleteTask(id string) (int, error)
	SearchTask(query, queryType string) ([]model.Task, error)
}
type Projects interface {
	GetProject() ([]model.Project, error)
	CreateProject(project *model.Project) (int, error)
	GetProjectById(id string) (model.Project, error)
	UpdateProject(project *model.Project, id string) (int, error)
	DeleteProject(id string) (int, error)
	SearchProject(query, queryType string) ([]model.Project, error)
	GetTasksForProject(projectId string) ([]model.Task, error)
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
