package service

import (
	_ "managep"
	"managep/pkg/repository"
)

type ProjectService struct {
	repo repository.Tasks
}

func NewProjectService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}
