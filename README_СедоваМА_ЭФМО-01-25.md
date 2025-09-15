в процессе

## Практическое занятие №1 Установка и настройка окружения Go.

### Выполнил: студент группы ЭФМО-01-25 Седова Мария Александровна.

### Цель: Развернуть рабочее окружение Go на Windows, создать минимальный HTTP-сервис на net/http, подключить и использовать внешнюю зависимость, собрать и проверить приложение.

### Задание:
1.	Установить Go и Git, проверить версии.
2.	Инициализировать модуль Go в новом проекте.
3.	Реализовать HTTP-сервер с маршрутами /hello (текст) и /user (JSON).
4.	Подключить внешнюю библиотеку (генерация UUID) и использовать её в /user.
5.	Запустить и проверить ответы curl/браузером.
6.	Собрать бинарник .exe и подготовить README и отчёт.

## Запуск и конфигурация

### Подготовка окружения и проверка версий

```
go version
git --version
```
<img width="522" height="124" alt="image" src="https://github.com/user-attachments/assets/c6fb79e6-d8c0-4440-8df7-aa4f3be918c6" />

### Модуль Go

```
go mod init example.com/helloapi
```
«Паспорт» проекта с именем модуля и версией Go.
<img width="537" height="219" alt="image" src="https://github.com/user-attachments/assets/8bb3487f-6e19-49b6-bea7-2ef38ae91288" />

### Структура проекта

<img width="212" height="221" alt="image" src="https://github.com/user-attachments/assets/39c90de3-a8d7-4330-b09b-d173abd53fb9" />

### Минимальный HTTP-сервер 
```
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	// Текстовый ответ
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	// Пока временный JSON (без UUID — добавим на шаге 4)
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(user{
			ID:   "temp",
			Name: "Gopher",
		})
	})

	addr := ":8080"
	log.Printf("Starting on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
```

### Подключение внешней зависимости и доработка /user
```
go get github.com/google/uuid@latest
go mod tidy
```
Обновляем код: добавляем в него библиотеку для импорта "github.com/google/uuid" и заменяем в функции для маршрута /user ID: "temp" на ID: uuid.NewString().
Появился файл "go.sum".

### Запуск сервера и быстрая проверка
```bash
go run ./cmd/server
```
<img width="635" height="63" alt="image" src="https://github.com/user-attachments/assets/8c2e834c-b5df-4d4b-a5e4-ed5e23b3408a" />

В другом окне PowerShell проверяем эндпоинты:
```
curl http://localhost:8080/hello
curl http://localhost:8080/user
```
1) <img width="833" height="528" alt="image" src="https://github.com/user-attachments/assets/6d6680a7-6e1f-4774-9e15-1dbb1718f79e" />
2) <img width="858" height="873" alt="image" src="https://github.com/user-attachments/assets/a71b8c00-d70b-40df-a0eb-652c350aca87" />



### Запуск на другом порту
Порт можно изменить через переменную окружения `APP_PORT`.

**Windows (PowerShell):**
```powershell
$env:APP_PORT="8081"
go run ./cmd/server
```

**Linux/macOS (Bash):**
```bash
export APP_PORT="8081"
go run ./cmd/server
```

### Примеры HTTP-запросов
После запуска сервер будет доступен по адресу `http://localhost:PORT`.

*   **GET /hello**
    ```bash
    curl http://localhost:8080/hello
    ```
    Ответ: `Hello, world!`

*   **GET /user**
    ```bash
    curl http://localhost:8080/user
    ```
    Ответ: `{"id":"a1b2c3d4-...", "name":"Gopher"}`

*   **GET /health** (бонусный эндпоинт)
    ```bash
    curl http://localhost:8080/health
    ```
    Ответ: `{"status":"ok", "time":"2023-10-25T15:04:05Z"}`

	Примечания по конфигурации (порт, переменные окружения).

3.	Отчётные материалы
	Скриншоты go version и ответов /hello и /user.
	Ссылка на репозиторий.
	
Итоговая проверка (чек-лист)
	Репозиторий клонируется, go build проходит без ошибок.
	go run ./cmd/server запускается, /hello и /user отвечают 200.
	В go.mod и go.sum зафиксированы зависимости; UUID реально генерируется.
	README содержит шаги запуска и примеры запросов.
	Код отформатирован (go fmt), базовая проверка go vet — без критики.
