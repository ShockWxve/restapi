package handlers

import (
	"context"
	"log"

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
	log.Println("GET request")

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
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}

		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	panic("unimplemented")
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) PatchTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	panic("unimplemented")
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	panic("unimplemented")
}

// func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("GET request")

// 	tasks, err := h.Service.ReadAllTasks()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("POST request")

// 	var task taskService.Task

// 	dec := json.NewDecoder(r.Body)
// 	if err := dec.Decode(&task); err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, e, http.StatusBadRequest)
// 		return
// 	}

// 	createdTask, err := h.Service.CreateTask(task)
// 	if err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, e, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(createdTask)
// }

// func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("PATCH request")

// 	vars := mux.Vars(r)
// 	idStr := vars["id"]

// 	id, err := strconv.ParseUint(idStr, 10, strconv.IntSize)
// 	if err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, "invalid id", http.StatusBadRequest)
// 		return
// 	}
// 	taskID := uint(id)

// 	var updateData struct {
// 		Task   string `json:"task,omitempty"`
// 		IsDone *bool  `json:"is_done,omitempty"`
// 	}

// 	// Декодируем тело запроса в структуру updateData
// 	dec := json.NewDecoder(r.Body)
// 	if err := dec.Decode(&updateData); err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, e, http.StatusBadRequest)
// 		return
// 	}

// 	updates := make(map[string]interface{})
// 	if updateData.Task != "" {
// 		updates["task"] = updateData.Task
// 	}
// 	if updateData.IsDone != nil {
// 		updates["is_done"] = *updateData.IsDone
// 	}

// 	// Проверяем, есть ли изменения
// 	if len(updates) == 0 {
// 		e := "empty fields"
// 		log.Println(e)
// 		http.Error(w, e, http.StatusBadRequest)
// 		return
// 	}

// 	updatedTask, err := h.Service.UpdateTaskByID(taskID, updates)
// 	if err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, e, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(updatedTask)
// }

// func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("DELETE request")

// 	vars := mux.Vars(r)
// 	idStr := vars["id"]

// 	id, err := strconv.ParseUint(idStr, 10, strconv.IntSize)
// 	if err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, "invalid id", http.StatusBadRequest)
// 		return
// 	}
// 	taskID := uint(id)

// 	err = h.Service.DeleteTaskByID(taskID)
// 	if err != nil {
// 		e := err.Error()
// 		log.Println(e)
// 		http.Error(w, e, http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
