# go-app-layout

# Организация домена в единой папке app

```
└── internal_with_app
    └── some_lib -- какая-нибудь библиотека  
    └── app      -- доменный уровень, тут располагаются доменные сущности и бизнес-логика, которая не зависит от деталей реализации внешних зависимотей
        └── storage -- подпакет app, уровень взаимодействия с хранилищами, другими словами адаптер домена к внешним зависимостям, которые отвечают за хранение данных
            ├── postgres -- конкретные реализации postgres репозиториев
                ├── registration_repository -- конкретная реализация postgres-репозитория, который отвечает за CRUD для app.Registration 
            └── memory   -- конкретные реализации in-memory репозиториев
                ├── registration_repository -- конкретная реализация in-memory репозитория, который отвечает за CRUD для app.Registration 
        └── flight.go
        └── flight_registrator.go
        └── flight_registrator_test.go
        └── passenger.go
        └── registration.go
    └── utils  
```


# Организация домена с распределением домена по разным пакетам

```
└── internal_with_models
    └── some_lib -- какая-нибудь библиотека 
    └── models   -- доменный уровень, тут располагаются доменные сущности
        └── flight.go
        └── passenger.go
        └── registration.go
    └── services -- доменный уровень, тут располагается бизнес-логика   
        └── flight_registrator.go
        └── flight_registrator_test.go
    └── storage  -- уровень взаимодействия с хранилищами, другими словами адаптер домена к внешним зависимостям, которые отвечают за хранение данных
        ├── postgres -- конкретные реализации postgres репозиториев 
            ├── registration_repository -- конкретная реализация postgres-репозитория, который отвечает за CRUD для app.Registration 
        └── memory   -- конкретные реализации in-memory репозиториев
    └── utils  
```
