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
