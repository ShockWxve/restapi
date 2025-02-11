package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// var task string

type RequestBody struct {
	Message string `json:"message,omitempty"`
}

// func getHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("GET request") // Лог о получении запроса

// 	// Возвращаем ответ
// 	if task == "" {
// 		fmt.Fprint(w, "hello!")
// 	} else {
// 		fmt.Fprintf(w, "hello, %v", task)
// 	}
// }

// func postHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println("POST request") // Лог о получении запроса

// 	requestBody := new(RequestBody)

// 	dec := json.NewDecoder(r.Body)
// 	if err := dec.Decode(requestBody); err != nil { // Десериализуем body
// 		e := fmt.Sprintln("Ошибка:", err) // Обрабатываем ошибки
// 		log.Print(e)
// 		fmt.Fprint(w, e)
// 	} else {
// 		fmt.Fprint(w, "OK!") // Возвращаем ответ
// 		// task = requestBody.Message // Присваиваем значение из body в глобальную переменную
// 	}
// }

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
	enc := json.NewEncoder(w)
	enc.Encode(tasks)
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
	successMessage := "{\n\t\"%v\" : \"Message created\"\n}"
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, successMessage, task.Task)
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
