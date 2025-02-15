package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// Передает в функцию task c типом Task из orm.go. Возвращает созданный Task и ошибку
	CreateTask(task Task) (Task, error)

	// Возвращает массив всех задач в БД и ошибку
	ReadAllTasks() ([]Task, error)

	// Передает id и Task. Возвращает обновленный Task и ошибку
	UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error)

	// Передает id для удаления, возвращает только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

// (r *taskRepository) привязывает данную функцию к нашему репозиторию

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) ReadAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error) {
	// Проверяем существование id
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
	task := new(Task)

	// Проверяем существование id
	if err := r.db.First(task, id).Error; err != nil {
		return err
	}

	// Удаляем запись из БД
	if err := r.db.Delete(task).Error; err != nil {
		return err
	}
	return nil
}
