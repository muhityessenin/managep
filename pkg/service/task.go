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
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask() ([]model.Task, error) {
	return s.repo.GetTask()
}

func (s *TaskService) GetTaskById(id string) (model.Task, error) {
	return s.repo.GetTaskById(id)
}
func (s *TaskService) UpdateTask(task *model.Task, id string) (int, error) {
	return s.repo.UpdateTask(task, id)
}

func (s *TaskService) DeleteTask(id string) (int, error) {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) SearchTask(query, queryType string) ([]model.Task, error) {
	return s.repo.SearchTask(query, queryType)
}
