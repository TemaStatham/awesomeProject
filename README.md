# Запуск

```bash
docker-compose up -d --build
go run ./cmd/api/main.go --config="./config/config.yaml"
```

# ЗАДАЧА
Развернуть сервис на Golang, Postgres, Clickhouse, Nats (альтернатива kafka), Redis
Описать модели данных и миграций
В миграциях Postgres
Проставить primary-key и индексы на указанные поля
При добавлении записи в таблицу устанавливать приоритет как макс приоритет в таблице +1. Приоритеты начинаются с 1
При накатке миграций добавить одну запись в Projects таблицу по умолчанию
id = serial
name = Первая запись
Реализовать CRUD методы на GET-POST-PATCH-DELETE данных в таблице GOODS в Postgres
При редактировании данных в Postgres ставить блокировку на чтение записи и оборачивать все в транзакцию. Валидируем поля при редактировании.
При редактировании данных в GOODS инвалидируем данные в REDIS
Если записи нет (проверяем на PATCH-DELETE), выдаем ошибку (статус 404)
code = 3
message = “errors.good.notFound“
details = {}
При GET запросе данных из Postgres кешировать данные в Redis на минуту. Пытаемся получить данные сперва из Redis, если их нет, идем в БД и кладем их в REDIS
При добавлении, редактировании или удалении записи в Postgres писать лог в Clickhouse через очередь Nats (альтернатива kafka). Логи писать пачками в Clickhouse
При обращении в БД использовать чистый SQL
Обернуть приложение в докер