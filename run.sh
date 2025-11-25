#!/bin/bash

echo "Проверка кода при помощи go vet"

go vet ./cmd
go vet ./internal/api
go vet ./internal/config
go vet ./internal/entities
go vet ./internal/provider
go vet ./internal/usecase

echo "Файлы проверены"

echo "Запуск тестов"

go test -v  ./internal/usecase/

echo "Успешный успех"

echo "Запуск сервера"

cd cmd
go run main.go