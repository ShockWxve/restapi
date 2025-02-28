package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shock_wave/restapi/internal/database"
	"github.com/shock_wave/restapi/internal/handlers"
	"github.com/shock_wave/restapi/internal/taskService"
	"github.com/shock_wave/restapi/internal/web/tasks"
)

func main() {
	database.InitDB()
	// Выключаем AutoMigrate при использовании makefile
	// database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	// router := mux.NewRouter()
	// router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	// router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	// router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	// router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	// log.Println("Сервер запущен")
	// http.ListenAndServe(":8080", router)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
