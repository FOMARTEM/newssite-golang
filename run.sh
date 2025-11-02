#!/bin/bash

echo "Проверка кода при помощи go vet"

go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/cmd
go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/internal/api
go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/internal/config
go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/internal/entities
go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/internal/provider
go vet /home/shahov/Documents/SFU/OPPO/newssite-golang/internal/usecase

echo "Файлы проверены"

echo "Запуск сервера"

cd cmd
go run main.go