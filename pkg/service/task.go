package service

import (
	_ "managep"
	"managep/pkg/repository"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}
