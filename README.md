# Avito Test Project
## v1.0.0

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

[Конфигурация для postman](https://pastebin.com/gmsnDngS)