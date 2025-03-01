# 🏠 SkillsRock API

Простое CRUD API для управления задачами (Tasks), написанное на **Go** с использованием **Fiber, PostgreSQL, SQLC**.

## 🚀 Технологии
- Go 1.21+
- Fiber (быстрый HTTP-фреймворк)
- PostgreSQL
- SQLC (автоматическая генерация Go-кода для SQL-запросов)
- Goose (миграции БД)
- pgx (Go-драйвер для PostgreSQL)

---

## 🔧 Установка и запуск

### 1️⃣ Клонирование репозитория
```sh
git clone https://github.com/HappyFreeman/SkillsRock.git
cd SkillsRock
```

### 2️⃣ Настройка `.env`
Создай файл `.env` и добавь настройки БД:
```
DB_URL=postgres://username:password@localhost:5432/database_name
```

### 3️⃣ Установка зависимостей
```sh
go mod tidy
```

### 4️⃣ Применение миграций
```sh
goose postgres $DB_URL up
```

### 5️⃣ Генерация SQL-кода (если изменял запросы)
```sh
sqlc generate
```

### 6️⃣ Запуск сервера
```sh
go run main.go
```
👉 Сервер запустится на `http://localhost:3000/api/v1/`

---

## 📌 API Эндпоинты

### ✅ Получить все задачи
```http
GET /api/v1/tasks
```

### ➕ Создать задачу
```http
POST /api/v1/tasks
Content-Type: application/json
```
```json
{
  "title": "Купить хлеб",
  "description": "Сходить в магазин",
  "status": "new"
}
```

### 🗒️ Обновить задачу
```http
PUT /api/v1/tasks/{id}
```

### ❌ Удалить задачу
```http
DELETE /api/v1/tasks/{id}
```

---

## 📂 Структура проекта
```
/SkillsRock
👉 /internal
   👉 /database   # Сгенерированные SQLC-запросы
👉 /sql
   👉 /schema     # SQL-миграции (Goose)
   👉 /queries    # SQL-запросы для SQLC
👉 main.go         # Точка входа
👉 task_handlers.go # Логика API
👉 go.mod          # Go зависимости
👉 README.md       # Описание проекта
```

---

## 🛠 Дополнительные команды
### 📌 Установка `goose` и `sqlc`
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

---

## 💡 Авторы
- **HappyFreeman** — разработка и архитектура проекта.

---

## 🎯 TODO (Дальнейшие улучшения)
- [ ] Добавить аутентификацию (JWT)
- [ ] Написать тесты для API
- [ ] Добавить Swagger-документацию

---

## 💜 Лицензия
Этот проект распространяется под лицензией MIT.

