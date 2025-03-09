package userService

import (
	"github.com/shockwxve/restapi/internal/taskService"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	ReadAllUsers() ([]User, error)
	UpdateUserByID(id uint, updates map[string]interface{}) (User, error)
	DeleteUserByID(id uint) error
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) ReadAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, updates map[string]interface{}) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	if err := r.db.Model(&user).Updates(updates).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
