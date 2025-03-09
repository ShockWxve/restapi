package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	ReadAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error)
	DeleteTaskByID(id uint) error
	ReadTasksByUserID(userID uint) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	return task, result.Error
}

func (r *taskRepository) ReadAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	if err := r.db.Model(&task).Updates(updates).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) ReadTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}
