package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shockwxve/restapi/internal/taskService"
	"github.com/shockwxve/restapi/internal/web/tasks"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (th *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := th.Service.ReadAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (th *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if taskRequest.UserId == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "user_id is required")
	}

	taskToCreate := taskService.Task{
		Task:   taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}

	createdTask, err := th.Service.CreateTask(taskToCreate, taskRequest.UserId)
	if err != nil {
		return nil, err
	}

	return tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}, nil
}

func (th *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := uint(request.Id)

	if request.Body.Task == nil && request.Body.IsDone == nil && request.Body.UserId == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "No updates provided")
	}

	updates := make(map[string]interface{})
	if request.Body.Task != nil {
		updates["task"] = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updates["is_done"] = *request.Body.IsDone
	}
	if request.Body.UserId != nil {
		updates["user_id"] = *request.Body.UserId
	}

	updatedTask, err := th.Service.UpdateTaskByID(id, updates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return nil, err
	}

	return tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}, nil
}

func (th *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := uint(request.Id)

	err := th.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (th *TaskHandler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userID := uint(request.Id)

	userTasks, err := th.Service.GetTasksByUserID(userID)
	if err != nil {
		return nil, err
	}

	var response tasks.GetUsersIdTasks200JSONResponse
	for _, task := range userTasks {
		response = append(response, tasks.Task{
			Id:     &task.ID,
			Task:   &task.Task,
			IsDone: &task.IsDone,
			UserId: &task.UserID,
		})
	}

	return response, nil
}
