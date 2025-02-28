package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shock_wave/restapi/internal/taskService"
	"github.com/shock_wave/restapi/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.ReadAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     tsk.ID,
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
		}

		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     createdTask.ID,
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
	}
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := uint(request.Id)
	updates := make(map[string]interface{})

	if request.Body.Task != nil {
		updates["task"] = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updates["is_done"] = *request.Body.IsDone
	}

	if len(updates) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "No updates provided")
	}

	updatedTask, err := h.Service.UpdateTaskByID(id, updates)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     updatedTask.ID,
		Task:   updatedTask.Task,
		IsDone: updatedTask.IsDone,
	}
	return response, nil
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
