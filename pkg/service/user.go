package service

import (
	_ "managep"
	"managep/pkg/model"
	"managep/pkg/repository"
	"net/http"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUser() ([]model.User, error) {
	res, err := u.repo.GetUser()
	if err != nil {
		return nil, err
	}
	for i := range res {
		res[i].RegistrationDate = res[i].RegistrationDate[:10]
	}
	return res, nil
}

func (u *UserService) GetUserById(id string) (model.User, error) {
	res, err := u.repo.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}
	res.RegistrationDate = res.RegistrationDate[:10]
	return res, nil
}

func (u *UserService) CreateUser(user *model.User) (int, error) {
	res, err := u.repo.CreateUser(user)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return res, nil
}

func (u *UserService) UpdateUser(user *model.User, id string) (int, error) {
	return u.repo.UpdateUser(user, id)
}

func (u *UserService) DeleteUser(id string) (int, error) {
	return u.repo.DeleteUser(id)
}

func (u *UserService) GetTasksForUser(id string) ([]model.Task, error) {
	res, err := u.repo.GetTasksForUser(id)
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

func (u *UserService) SearchUser(query, queryType string) (model.User, error) {
	res, err := u.repo.SearchUser(query, queryType)
	if err != nil {
		return model.User{}, err
	}
	res.RegistrationDate = res.RegistrationDate[:10]
	return res, nil
}
