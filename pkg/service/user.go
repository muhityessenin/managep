package service

import (
	_ "managep"
	"managep/pkg/repository"
)

type UserService struct {
	repo repository.Tasks
}

func NewUserService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}
