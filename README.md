# Web API Server

### Структура
```bash
├───cmd
├───config
├───internal
│   ├───app
│   ├───models
│   ├───repository
│   │   ├───revenue
│   │   └───user
│   ├───services
│   └───transport
│       └───rest
│           ├───revenue
│           └───user
└───pkg
    ├───cache
    ├───client
    │   └───postgres
    ├───logger
    └───utils
```

### Unit тесты
>   Для запуска тестов:

```bash
make test
```

>   Для проверки покрытия тестов:

```bash
make view_test
```

>   Для проверки покрытия
```bash
go test -v ./ -coverprofile=filename
go tool cover -html=filename
```