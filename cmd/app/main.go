package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shockwxve/restapi/internal/database"
	"github.com/shockwxve/restapi/internal/handlers"
	"github.com/shockwxve/restapi/internal/taskService"
	"github.com/shockwxve/restapi/internal/userService"
	"github.com/shockwxve/restapi/internal/web/tasks"
	"github.com/shockwxve/restapi/internal/web/users"
)

func main() {
	// Инициализируем БД
	database.InitDB()

	// Инициализируем репозитории
	tasksRepo := taskService.NewRepository(database.DB)
	usersRepo := userService.NewRepository(database.DB)

	// Инициализируем сервисы
	tasksService := taskService.NewService(tasksRepo)
	usersService := userService.NewService(usersRepo)

	// Инициализируем хендлеры
	tasksHandler := handlers.NewTaskHandler(tasksService)
	usersHandler := handlers.NewUserHandler(usersService)

	// Создаём Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем хендлеры для задач
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Регистрируем хендлеры для пользователей
	usersStrictHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запускаем сервер на 8080 порту
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
