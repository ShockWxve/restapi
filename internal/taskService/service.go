package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) ReadAllTasks() ([]Task, error) {
	return s.repo.ReadAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, updates map[string]interface{}) (Task, error) {
	return s.repo.UpdateTaskByID(id, updates)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
