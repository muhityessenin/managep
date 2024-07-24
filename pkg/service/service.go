package service

import (
	_ "managep"
	"managep/pkg/repository"
)

type Users interface {
}
type Tasks interface {
}
type Projects interface {
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
