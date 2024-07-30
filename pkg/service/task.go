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
	res, err := s.repo.GetTask()
	if err != nil {
		return make([]model.Task, 0), err
	}
	for i := range res {
		res[i].CreatedAt = res[i].CreatedAt[:10]
		if len(res[i].FinishedAt) >= 10 {
			res[i].FinishedAt = res[i].FinishedAt[:10]
		}
	}
	return res, nil
}

func (s *TaskService) GetTaskById(id string) (model.Task, error) {
	res, err := s.repo.GetTaskById(id)
	if err != nil {
		return model.Task{}, err
	}
	res.CreatedAt = res.CreatedAt[:10]
	if len(res.FinishedAt) >= 10 {
		res.FinishedAt = res.FinishedAt[:10]
	}
	return res, nil
}
func (s *TaskService) UpdateTask(task *model.Task, id string) (int, error) {
	return s.repo.UpdateTask(task, id)
}

func (s *TaskService) DeleteTask(id string) (int, error) {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) SearchTask(query, queryType string) ([]model.Task, error) {
	res, err := s.repo.SearchTask(query, queryType)
	if err != nil {
		return make([]model.Task, 0), err
	}
	for i := range res {
		res[i].CreatedAt = res[i].CreatedAt[:10]
		if len(res[i].FinishedAt) >= 10 {
			res[i].FinishedAt = res[i].FinishedAt[:10]
		}
	}
	return res, nil
}
