#!/bin/bash

set -e  # Скрипт завершится, если какая-то команда вернёт ненулевой код

echo "Проверка кода при помощи go vet"

go vet ./cmd
go vet ./internal/api
go vet ./internal/config
go vet ./internal/entities
go vet ./internal/provider
go vet ./internal/usecase

echo "Файлы проверены"
echo "Запуск сервера"

cd cmd || exit 1

# Запускаем сервер в фоне
go run main.go &
SERVER_PID=$!

echo "Сервер запущен (PID: $SERVER_PID)"
echo "Введите 'stop' чтобы остановить сервер"

# Обработка Ctrl+C (SIGINT) и завершение скрипта
trap "echo 'Остановка сервера...'; kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null; echo 'Сервер остановлен'; exit 0" SIGINT

# Цикл ожидания ввода
while true; do
  read -r cmd
  if [[ "$cmd" == "stop" ]]; then
    echo "Остановка сервера..."
    kill $SERVER_PID 2>/dev/null
    wait $SERVER_PID 2>/dev/null
    echo "Сервер остановлен"
    break
  else
    echo "Неизвестная команда: $cmd (введите 'stop' для остановки)"
  fi
done
