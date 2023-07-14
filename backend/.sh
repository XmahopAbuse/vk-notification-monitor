#!/bin/bash

# Ожидание доступности базы данных
until nc -z db 5432; do
  echo "Waiting for the database to be available..."
  sleep 1
done

# Запуск приложения
echo "Starting the application..."
exec "/app/main"