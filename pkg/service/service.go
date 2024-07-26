package service

import (
	_ "managep"
	"managep/pkg/model"
	"managep/pkg/repository"
)

type Users interface {
	GetUser() ([]model.User, error)
	CreateUser(user *model.User) (int, error)
	GetUserById(id string) (model.User, error)
	UpdateUser(user *model.User, id string) (int, error)
	DeleteUser(id string) (int, error)
	GetTasksForUser(id string) ([]model.Task, error)
	SearchUser(query, queryType string) ([]model.User, error)
}
type Tasks interface {
	GetTask() ([]model.Task, error)
	CreateTask(task *model.Task) (int, error)
	GetTaskById(id string) (model.Task, error)
	UpdateTask(task *model.Task, id string) (int, error)
	DeleteTask(id string) (int, error)
}
type Projects interface {
	CreateProject(project *model.Project) (int, error)
}
type Service struct {
	UserService    Users
	TaskService    Tasks
	ProjectService Projects
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService:    NewUserService(repository.UserService),
		TaskService:    NewTaskService(repository.TaskService),
		ProjectService: NewProjectService(repository.ProjectService),
	}
}
