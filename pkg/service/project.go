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
	return s.repo.GetProject()
}

func (s *ProjectService) GetProjectById(id string) (model.Project, error) {
	return s.repo.GetProjectById(id)
}

func (s *ProjectService) UpdateProject(project *model.Project, id string) (int, error) {
	return s.repo.UpdateProject(project, id)
}

func (s *ProjectService) DeleteProject(id string) (int, error) {
	return s.repo.DeleteProject(id)
}

func (s *ProjectService) SearchProject(query, queryType string) ([]model.Project, error) {
	return s.repo.SearchProject(query, queryType)
}

func (s *ProjectService) GetTasksForProject(projectId string) ([]model.Task, error) {
	return s.repo.GetTasksForProject(projectId)
}
