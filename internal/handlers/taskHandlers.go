package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shock_wave/restapi/internal/taskService"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET request")

	tasks, err := h.Service.ReadAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request")

	var task taskService.Task

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&task); err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PATCH request")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, strconv.IntSize)
	if err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	taskID := uint(id)

	var updateData struct {
		Task   string `json:"task,omitempty"`
		IsDone *bool  `json:"is_done,omitempty"`
	}

	// Декодируем тело запроса в структуру updateData
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&updateData); err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{})
	if updateData.Task != "" {
		updates["task"] = updateData.Task
	}
	if updateData.IsDone != nil {
		updates["is_done"] = *updateData.IsDone
	}

	// Проверяем, есть ли изменения
	if len(updates) == 0 {
		e := "empty fields"
		log.Println(e)
		http.Error(w, e, http.StatusBadRequest)
		return
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, updates)
	if err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE request")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseUint(idStr, 10, strconv.IntSize)
	if err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	taskID := uint(id)

	// var deleteData struct {
	// 	ID uint `json:"id"`
	// }

	// dec := json.NewDecoder(r.Body)
	// if err := dec.Decode(&deleteData); err != nil {
	// 	e := err.Error()
	// 	log.Println(e)
	// 	http.Error(w, e, http.StatusBadRequest)
	// 	return
	// }

	err = h.Service.DeleteTaskByID(taskID)
	if err != nil {
		e := err.Error()
		log.Println(e)
		http.Error(w, e, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
