# todo-list

## Описание

Данный проект представляет собой приложение для создания списка задач 
(todo-list).

## Структура проекта

```text
├── cmd
│   └── server
│       └── main.go // точка входа в приложение
│
├── configs
│   └── config.yml // файл с конфигами
│
├── docs // swagger документация
│
├── internal
│   ├── app // слой бизнес-логики
│   │   ├── mocks
│   │   ├── valid // пакет для валидации полей
│   │   ├── app.go // реализация интерфейса приложения
│   │   ├── app_interface.go // интерфейс приложения
│   │   └── app_test.go
│   │
│   ├── model // слой сущностей (entities)
│   │   ├── date.go
│   │   ├── errs.go
│   │   └── todo_task.go // структура задачи
│   │
│   ├── ports // сетевой слой (infrastructure)
│   │   └── httpserver // rest-сервер
│   │       ├── handlers.go
│   │       ├── presenters.go
│   │       ├── responses.go
│   │       ├── router.go
│   │       └── server.go
│   │
│   └── repo // хранилище задач
│       └── repo.go
│
├── migrations
│   └── task_repo_init.sql // скрипт для конфигурации task_repo
│
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
└── go.sum

```

## Бизнес-логика

Реализован CRUDL для задач, состоящих из заголовка, описания, запланированной 
даты и статуса (выполнено/не выполнено). Перед добавлением/обновлением 
происходит валидация задачи:

* Заголовок не пустой и его длина не больше 100 байтов
* Описание не больше 500 байтов
* Запланированная дата не раньше даты на момент добавления/обновления

Помимо поиска по id задачи реализован регистронезависимый поиск по вхождению 
искомого текста в заголовок/описание задачи.

Реализовано получение списка задач с фильтром по статусу и пагинацией, либо с 
фильтром по дате и статусу.

## Используемые технологии

* go 1.21
* PostgreSQL
* Docker
* Gin Web Framework
* Swagger

## Запуск приложения

### С помощью Docker

```shell
docker-compose up
```

### Локально

Самостоятельно сконфигурировать БД PostgreSQL, используя [файл миграции](https://github.com/papey08/todo-list/blob/master/migrations/task_repo_init.sql), 
заменить в файле [**config.yml**](https://github.com/papey08/todo-list/blob/master/configs/config.yml) 
конфигурационные данные на свои, после чего выполнить команды:

```shell
go mod download
go run cmd/server/main.go
```

### Запуск тестов

```shell
go mod tidy
go test -v -race ./... ./...
```

## Формат запросов

Swagger-документация доступна по адресу http://localhost:8080/todo-list/api/swagger/index.html 
либо в файле [***swagger.json***](https://github.com/papey08/todo-list/blob/master/docs/swagger.json)

### Добавление новой задачи

* Метод: `POST`
* Эндпоинт: `http://localhost:8080/todo-list/api/task`
* Формат тела запроса:

```json
{
    "title": "title",
    "description": "description",
    "planning_date": {
        "year": 2024,
        "month": 1,
        "day": 1
    },
    "status": false
}

```

* Формат ответа:

```json
{
    "data": {
        "id": 1,
        "title": "title",
        "description": "description",
        "planning_date": {
            "year": 2024,
            "month": 1,
            "day": 1
        },
        "status": false
    },
    "error": null
}
```

### Получение задачи по id

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/todo-list/api/task/1`
* Формат ответа:

```json
{
    "data": {
        "id": 1,
        "title": "title",
        "description": "description",
        "planning_date": {
            "year": 2024,
            "month": 1,
            "day": 1
        },
        "status": false
    },
    "error": null
}
```

### Поиск задачи по тексту заголовка или описания

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/todo-list/api/task`
* Формат тела запроса:

```json
{
    "text": "title"
}
```

* Формат ответа:

```json
{
    "data": [
        {
            "id": 1,
            "title": "title",
            "description": "description",
            "planning_date": {
                "year": 2024,
                "month": 1,
                "day": 1
            },
            "status": false
        }
    ],
    "error": null
}
```

### Обновление задачи

* Метод: `PUT`
* Эндпоинт: `http://localhost:8080/todo-list/api/task/1`
* Формат тела запроса:

```json
{
    "title": "title",
    "description": "description",
    "planning_date": {
        "year": 2024,
        "month": 1,
        "day": 1
    },
    "status": true
}
```

* Формат ответа:

```json
{
    "data": {
        "id": 1,
        "title": "title",
        "description": "description",
        "planning_date": {
            "year": 2024,
            "month": 1,
            "day": 1
        },
        "status": true
    },
    "error": null
}
```

### Удаление задачи

* Метод: `DELETE`
* Эндпоинт: `http://localhost:8080/todo-list/api/task/1`
* Формат ответа:

```json
{
    "data": null,
    "error": null
}
```

### Получение списка задач с фильтром по статусу и пагинацией

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/todo-list/api/task/by_status`
* Формат тела запроса:

```json
{
    "status": true,
    "offset": 0,
    "limit": 1
}
```

* Формат ответа:

```json
{
    "data": [
        {
            "id": 1,
            "title": "title",
            "description": "description",
            "planning_date": {
                "year": 2024,
                "month": 1,
                "day": 1
            },
            "status": true
        }
    ],
    "error": null
}
```

### Получение списка задач с фильтром по дате и статусу

* Метод: `GET`
* Эндпоинт: `http://localhost:8080/todo-list/api/task/by_date`
* Формат тела запроса:

```json
{
  "planning_date": {
    "year": 2024,
    "month": 1,
    "day": 1
  },
  "status": true
}
```

* Формат тела ответа:

```json
{
    "data": [
        {
            "id": 1,
            "title": "title",
            "description": "description",
            "planning_date": {
                "year": 2024,
                "month": 1,
                "day": 1
            },
            "status": true
        }
    ],
    "error": null
}
```
