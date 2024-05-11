### Цель
> Разработка веб-приложения для сбора и хранения данных о пациентах и предоставления этих данных для дальнейших исследований

### Процесс разработки

1. [x] **Проектирование базы данных:** Начните с проектирования структуры базы данных и определения связей между таблицами.
2. [x] **Создание шаблонов:** Разработайте шаблоны как основу для логики приложения.
3. [x] **Реализация маршрутов:** Создайте маршруты для обработки входящих запросов.
4. [x] **Взаимодействие с базой данных:** Реализуйте запросы и взаимодействие с базой данных.
5. [x] **Интеграция:** Свяжите маршруты с взаимодействием с базой данных для обеспечения бесперебойной работы.
6. [] **Исправление ошибок:** Идентифицируйте и исправьте любые ошибки или проблемы в логике приложения.
7. [] **Тестирование:** Напишите комплексные тесты, чтобы обеспечить надежность приложения.
8. [] **Документация:** Задокументируйте приложение, конечные точки API и любую другую актуальную информацию.

### Запуск приложения

Чтобы запустить приложение локально, выполните следующие шаги:

1. Клонируйте репозиторий на ваш локальный компьютер.
2. Перейдите в директорию проекта.
3. Установите зависимости, используя команду `go mod download`.
4. Настройте базу данных в соответствии с инструкциями в файле `README.md`.
5. Запустите приложение с помощью команды `go run cmd/app/main.go`.
6. Получите доступ к приложению в вашем веб-браузере по адресу `http://localhost:8080`.

### Docker

Приложение также можно запустить с использованием Docker. Используйте следующие команды:

1. Соберите образ Docker: `docker build -t <имя_образа> .`
2. Запустите контейнер Docker: `docker run -p 8080:8080 <имя_образа>`

### Документация по API

- [Документация Swagger API](https://app.swaggerhub.com/apis/DANIILBAKHLANOV/oncomarker-api/1.0-oas3): Изучите конечные точки API и их описания.

### Дополнительные замечания

- Не забудьте настроить переменные окружения, такие как учетные данные базы данных и ключи API, перед запуском приложения.
- Отслеживайте журналы приложения на наличие ошибок или предупреждений во время выполнения.
- Регулярно обновляйте зависимости и пересматривайте кодовую базу на наличие уязвимостей безопасности или устаревших практик.