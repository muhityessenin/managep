package service

import (
	_ "managep"
	"managep/pkg/model"
	"managep/pkg/repository"
)

type ProjectService struct {
	repo repository.Projects
}

func NewProjectService(repo repository.Projects) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(project *model.Project) (int, error) {
	return s.repo.CreateProject(project)
}

func (s *ProjectService) GetProject() ([]model.Project, error) {
	result, err := s.repo.GetProject()
	if err != nil {
		return make([]model.Project, 0), err
	}
	for i := range result {
		result[i].StartDate = result[i].StartDate[:10]
		if len(result[i].FinishDate) >= 10 {
			result[i].FinishDate = result[i].FinishDate[:10]
		}
	}
	return result, nil
}

func (s *ProjectService) GetProjectById(id string) (model.Project, error) {
	result, err := s.repo.GetProjectById(id)
	if err != nil {
		return model.Project{}, err
	}
	result.StartDate = result.StartDate[:10]
	if len(result.FinishDate) >= 10 {
		result.FinishDate = result.FinishDate[:10]
	}
	return result, nil
}

func (s *ProjectService) UpdateProject(project *model.Project, id string) (int, error) {
	return s.repo.UpdateProject(project, id)
}

func (s *ProjectService) DeleteProject(id string) (int, error) {
	return s.repo.DeleteProject(id)
}

func (s *ProjectService) SearchProject(query, queryType string) ([]model.Project, error) {
	result, err := s.repo.SearchProject(query, queryType)
	if err != nil {
		return make([]model.Project, 0), err
	}
	for i := range result {
		result[i].StartDate = result[i].StartDate[:10]
		if len(result[i].FinishDate) >= 10 {
			result[i].FinishDate = result[i].FinishDate[:10]
		}
	}
	return result, nil
}

func (s *ProjectService) GetTasksForProject(projectId string) ([]model.Task, error) {
	result, err := s.repo.GetTasksForProject(projectId)
	if err != nil {
		return make([]model.Task, 0), err
	}
	for i := range result {
		result[i].CreatedAt = result[i].CreatedAt[:10]
		if len(result[i].FinishedAt) >= 10 {
			result[i].FinishedAt = result[i].FinishedAt[:10]
		}
	}
	return result, nil
}
