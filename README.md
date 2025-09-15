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

<img width="858" height="318" alt="image" src="https://github.com/user-attachments/assets/dee1008f-3084-4572-96ef-697cd18c8d8a" />

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

	// Пока временный JSON
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
```
go run ./cmd/server
```
<img width="649" height="62" alt="image" src="https://github.com/user-attachments/assets/346d15e1-4e34-4e33-b048-8bcd327cb8f8" />

В другом окне PowerShell проверяем эндпоинты:
```
curl http://localhost:8080/hello
curl http://localhost:8080/user
```
1) <img width="1193" height="733" alt="image" src="https://github.com/user-attachments/assets/dd67abad-57fc-41be-9187-0112eb96b1f1" />

2) <img width="1705" height="814" alt="image" src="https://github.com/user-attachments/assets/a03760dc-a32e-45f4-96a2-fed758180b14" />

### Сборка бинарника 
```
go build -o helloapi.exe ./cmd/server
.\helloapi.exe
```
### Код-стайл и базовые проверки
```
go fmt ./...
go vet ./...
```

<img width="689" height="133" alt="image" src="https://github.com/user-attachments/assets/30d758bb-2aa4-40e6-8695-2dcfa8289ce6" />

### Запуск на другом порту
Порт можно изменить через переменную окружения `APP_PORT`.
```
$env:APP_PORT="8081"
go run ./cmd/server
```
<img width="847" height="104" alt="image" src="https://github.com/user-attachments/assets/bbc0b4d2-d054-4672-b7c6-9c87c15a3693" />

<img width="1241" height="729" alt="image" src="https://github.com/user-attachments/assets/a82dd401-e0ea-4809-b998-54113989474e" />

### Бонусный эндпоинт

    ```
    curl http://localhost:8080/health
    ```
<img width="1572" height="751" alt="image" src="https://github.com/user-attachments/assets/9895142b-88a1-4f65-a180-8e2ebaec8697" />

