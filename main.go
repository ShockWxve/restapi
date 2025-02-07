package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type RequestBody struct {
	Message string `json:"message,omitempty"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET request") // Лог о получении get-запроса

	// Возвращаем ответ
	if task == "" {
		fmt.Fprint(w, "hello!")
	} else {
		fmt.Fprintf(w, "hello, %v", task)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("POST request") // Лог о получении запроса

	requestBody := new(RequestBody)

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(requestBody); err != nil { // Десериализуем body
		e := fmt.Sprintln("Ошибка:", err) // Обрабатываем ошибки
		log.Print(e)
		fmt.Fprint(w, e)

	} else {
		fmt.Fprint(w, "OK!")       // Возвращаем ответ
		task = requestBody.Message // Присваиваем значение из body в глобальную переменную
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", getHandler).Methods("GET")
	router.HandleFunc("/api/hello", postHandler).Methods("POST")

	port := ":8080"
	http.ListenAndServe(":8080", router)
	log.Println("Запущен сервер на", port)
}
