# Avito Test Project
## v2.0.0

**Как собрать проект:**

```
docker-compose up
```

Для первого запуска также потребуется сделать миграцию с помощью утилиты [migrate](https://github.com/golang-migrate/migrate)
```
migrate -path ./src/schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up 
```

Также при первом запуске контейнер **avito-test** может упасть с ошибкой.

fix: Просто перезапустить отдельно контейнер после того, как поднимется БД

```
docker-compose up avito-test
```

**Примеры запросов:**

[Postman](https://www.postman.com/lunar-module-astronomer-11061739/workspace/avito-test-task/request/24172792-2a76305b-9e4b-4c5d-90d9-23743f0a02ae)

**Детали реализации:**

1. В базе на прямую не хранится баланс пользоваетеля, хранятся лишь транзакции. Таким образом баланс не будет расходиться с реальностью
2. Резерв денег осуществляется засчет отдельной таблички, от туда же происходят возвраты и фиксация.