package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shock_wave/restapi/internal/database"
	"github.com/shock_wave/restapi/internal/handlers"
	"github.com/shock_wave/restapi/internal/taskService"
)

func main() {
	database.InitDB()
	// Выключаем AutoMigrate при использовании makefile
	// database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	log.Println("Сервер запущен")
	http.ListenAndServe(":8080", router)
}
