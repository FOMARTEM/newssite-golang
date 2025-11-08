#!/bin/bash

echo "Проверка кода при помощи go vet"

go vet ./cmd
go vet ./internal/api
go vet ./internal/config
go vet ./internal/entities
go vet ./internal/provider
go vet ./internal/usecase

echo "Файлы проверены"

echo "Запуск сервера"

cd cmd || exit

# Запускаем сервер в фоне
go run main.go &
SERVER_PID=$!

echo "Сервер запущен (PID: $SERVER_PID)"
echo "Введите 'stop' чтобы остановить сервер"

# Цикл ожидания команды stop
while true; do
  read -r cmd
  if [[ "$cmd" == "stop" ]]; then
    echo "Остановка сервера..."
    kill $SERVER_PID
    wait $SERVER_PID 2>/dev/null
    echo "Сервер остановлен"
    break
  else
    echo "Неизвестная команда: $cmd (введите 'stop' для остановки)"
  fi
done
