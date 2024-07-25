package service

import (
	_ "managep"
	"managep/pkg/model"
	"managep/pkg/repository"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *model.Task) (int, error) {
	return 200, nil
}
