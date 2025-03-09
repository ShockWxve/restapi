# Makefile

# Переменные
DB_DSN := "postgres://postgres:pwd1111!@127.0.0.1:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
# Запуск: make migrate-new NAME=create_something
migrate-new: 
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение всех миграций
migrate:
	$(MIGRATE) up

# Откат всех миграций
migrate-down:
	$(MIGRATE) down

# Применить только одну миграцию
migrate-up-one:
	$(MIGRATE) up 1

# Откатить последнюю миграцию
migrate-down-last:
	$(MIGRATE) down 1

# Запустить приложение
run:
	go run cmd/app/main.go

# Генерация кода OpenAPI (пример разделения по tags)
gen-tasks:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

# Линтер
lint:
	golangci-lint run --out-format=colored-line-number
