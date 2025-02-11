package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	Message string `json:"message,omitempty"`
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
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

func CreateMessage(w http.ResponseWriter, r *http.Request) {
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
	result := DB.Create(&task)
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

func main() {
	// Вызываем БД
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()

	// Список маршрутов
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")

	port := ":8080"
	log.Println("Запущен сервер на", port)
	http.ListenAndServe(port, router)
}
