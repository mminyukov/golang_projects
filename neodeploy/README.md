> :warning: Владелец репозитория **не является разработчиком**, а представленный код - любительский. Сам проект является попыткой разобраться в языке, чтобы решить конкретную задачу.

## Описание
Проект предназначен для локального деплоя в зависимости от настроек и используемого приложения.
Для запуска необходим файл **settingsdeploy.json**
Пример содержимого файла:
```json
{
  "cce":
    {
      "stand_name": "default",
      "target_service_name": "JobServer.dll",
      "recreate_db": false,
      "connection_string": "Host=test-db;Port=5432;Database=CCI;Username=postgres;Password=postgres;",
      "main_zip_name": "job-server.zip"
    },
  "oak":
    {
      "stand_name": "default",
      "port_site": "8000",
      "target_service_name": "dummy",
      "connection_string_user": "Host=test-db;Port=5432;Database=demo_db",
      "connection_string_db": "Host=test-db;Port=5432;Database=demo_db;Username=postgres;Password=postgres;",
      "connection_string_hangfire": "Host=test-db;Port=5432;Database=demo_db;Username=hangfire;Password=hangfire;",
      "connection_string_cci": "Host=test-db;Port=5432;Username=hangfire;Password=hangfire;Database=CCI;",
      "ss_url": "http://localhost:8000/",
      "target_service_name": "dummy",
      "dbupdater_zip_name": "dummy-dbupdater.zip",
      "main_zip_name": "job-server.zip",
      "use_scheduler": false,
      "port_scheduler": "6001",
      "target_scheduler_name": "dummy",
      "connection_string_scheduler": "http://127.0.0.1:6001/",
      "scheduler_zip_name": "dummy-web.zip",
      "recreate_db": false,
      "install_cce": false
    }
}
```

Для запуска обязателен один параметр (наименование проекта), например:
```bash
go run neodeploy.go cce
```
