package userService

import "github.com/shockwxve/restapi/internal/taskService"

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) ReadAllUsers() ([]User, error) {
	return s.repo.ReadAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, updates map[string]interface{}) (User, error) {
	return s.repo.UpdateUserByID(id, updates)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return s.repo.GetTasksForUser(userID)
}
