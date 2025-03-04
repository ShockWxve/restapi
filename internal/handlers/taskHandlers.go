package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shock_wave/restapi/internal/taskService"
	"github.com/shock_wave/restapi/internal/web/tasks"
	"gorm.io/gorm"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// // Для вывода полей в нужном порядке
// type CustomTaskResponse struct {
// 	Id     uint   `json:"id"`
// 	Task   string `json:"task"`
// 	IsDone bool   `json:"is_done"`
// }

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.ReadAllTasks()
	if err != nil {
		return nil, err
	}

	var response []tasks.Task

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		response = append(response, tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		})
	}
	return tasks.GetTasks200JSONResponse(response), nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.Task{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return tasks.PostTasks201JSONResponse(response), nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := uint(request.Id)

	// Сразу проверяем, есть ли обновления
	if request.Body.Task == nil && request.Body.IsDone == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "No updates provided")
	}

	// Формируем мапу обновлений
	updates := make(map[string]interface{})
	// Забираем из тела только непустые значения
	if request.Body.Task != nil {
		updates["task"] = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updates["is_done"] = *request.Body.IsDone
	}

	// Обновляем задачу
	updatedTask, err := h.Service.UpdateTaskByID(id, updates)
	switch err {
	case nil:
		response := tasks.Task{
			Id:     &updatedTask.ID,
			Task:   &updatedTask.Task,
			IsDone: &updatedTask.IsDone,
		}
		return tasks.PatchTasksId200JSONResponse(response), nil
	case gorm.ErrRecordNotFound:
		return nil, echo.NewHTTPError(http.StatusNotFound, "Task not found")
	default:
		return nil, err
	}

}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := uint(request.Id)

	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
