# calc_go
задания по выполнению calc_go 
Простой веб-сервис для вычисления арифметических выражений

Описание
Этот проект реализует веб-сервис для вычисления арифметических выражений, переданных пользователем через HTTP-запрос. Сервис принимает входные данные в формате JSON, обрабатывает их и возвращает результат вычислений в формате JSON.

Структура проекта
- `cmd/` — точка входа приложения, содержащая файл для запуска сервера.
- `internal/` — внутренняя логика и модули приложения, включая обработку арифметических выражений.
- `pkg/` — вспомогательные пакеты и утилиты, используемые в проекте.

Запуск сервиса
Для запуска сервиса выполните следующие действия:
1. Установите Go. Убедитесь, что у вас установлена версия Go 1.18 или выше. Вы можете скачать и установить Go с [официального сайта](https://golang.org/dl/).
2. Склонируйте проект с GitHub:

   ```bash
   git clone https://github.com/your-username/calc_service.git

*Замените your-username на ваше имя пользователя на GitHub.
Перейдите в папку проекта:

cd calc_service

Запустите сервер с помощью команды:

go run ./cmd/main.go

После успешного запуска сервиса вы увидите сообщение в консоли:

Starting server on :8080
Сервис будет доступен по адресу: http://localhost:8080/api/v1/calculate.

Формат запроса
Запрос должен содержать JSON объект со следующей структурой:

{
    "expression": "выражение, которое ввёл пользователь"
}

*Где "expression" — это строка, представляющая собой арифметическое выражение.

Примеры запросов
Успешный запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Ожидаемый ответ:

{
    "result": "6"
}

Запрос с ошибкой 422 (неверное выражение):

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2++2*2"
}'

Ожидаемый ответ:

{
    "error": "Expression is not valid"
}

Запрос с ошибкой 500 (внутренняя ошибка сервера):

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "10/0"
}'

Ожидаемый ответ:

{
    "error": "Internal server error"
}

Тестирование
Для запуска тестов выполните следующую команду в терминале:

go test ./...

Это выполнит все тесты из директории test/ и проверит работоспособность вашего приложения.

Примечания
Для работы API требуется установленный Go (версия 1.18 и выше).
Все зависимости проекта управляются через go mod. Убедитесь, что в корне проекта находятся файлы go.mod и go.sum.
