package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("GET request")

	var tasks []Task

	// Получаем все записи из БД
	result := DB.Find(&tasks)
	if result.Error != nil {
		e := fmt.Sprintln("Ошибка чтения из БД:", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Возвращаем успешный ответ
	w.Header().Set("Content-Type", "application/json") // Устанавливаем формат
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request") // Лог о получении запроса

	task := new(Task)

	// Декодируем тело запроса в структуру Task
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(task); err != nil {
		e := fmt.Sprintln("Ошибка декодирования JSON:", err) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Сохраняем задачу в базу данных
	result := DB.Create(task)
	if result.Error != nil {
		e := fmt.Sprintln("Ошибка записи в БД:", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Возвращаем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	log.Println("PATCH request")

	// Определяем структуру для декодирования body
	var updateData struct {
		ID     uint   `json:"id"`
		Task   string `json:"task,omitempty"`
		IsDone *bool  `json:"is_done,omitempty"`
	}

	// Декодируем тело запроса в структуру updateData
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&updateData); err != nil {
		e := fmt.Sprintln("Ошибка декодирования JSON:", err) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	task := new(Task)

	// Проверяем существование id
	result := DB.First(task, updateData.ID)
	if result.Error != nil {
		e := fmt.Sprintln("Таск не найден.", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Обновляем только переданные поля
	updates := map[string]interface{}{}
	if updateData.Task != "" {
		updates["task"] = updateData.Task
	} // можно лучше переписать if-ы
	if updateData.IsDone != nil {
		updates["is_done"] = updateData.IsDone
	}

	if result := DB.Model(task).Updates(updates); result.Error != nil {
		e := fmt.Sprintln("Ошибка обновления записи:", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Возвращаем обновлённую задачу
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE request")

	var deleteData struct {
		ID uint `json:"id"`
	}

	// Декодируем body
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&deleteData); err != nil {
		e := fmt.Sprintln("Ошибка декодирования JSON:", err) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	task := new(Task)

	// Проверяем существование id
	if result := DB.First(task, deleteData.ID); result.Error != nil {
		e := fmt.Sprintln("Таск не найден.", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Удаляем запись из БД
	if result := DB.Delete(task); result.Error != nil {
		e := fmt.Sprintln("Ошибка удаления таска:", result.Error) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)
		return
	}

	// Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	// Вызываем БД
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	// Список маршрутов
	router.HandleFunc("/api/messages", GetTasks).Methods("GET")
	router.HandleFunc("/api/messages", CreateTask).Methods("POST")
	router.HandleFunc("/api/messages", UpdateTask).Methods("PATCH")
	router.HandleFunc("/api/messages", DeleteTask).Methods("DELETE")

	port := ":8080"
	log.Println("Запущен сервер на", port)
	http.ListenAndServe(port, router)
}
