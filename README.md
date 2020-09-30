# go-app-layout

# Организация домена в единой папке app

```
└── internal_with_app
    └── some_client   
    └── app -- доменный уровень, тут располагаются доменные сущности и бизнес-логика, которая не зависит от деталей реализации внешних зависимотей
        └── storage -- подпакет app, уровень взаимодействия с хранилищами, другими словами адаптер домена к внешним зависимостям, которые отвечают за хранение данных
            ├── postgres -- конкретная реализация 
                ├── registration_repository -- конкретная реализация postgres-репозитория, который отвечает за CRUD для app.Registration 
            └── memory 
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
    └── some_client   
    └── models -- доменный уровень, тут располагаются доменные сущности и бизнес-логика, которая не зависит от деталей реализации внешних зависимотей
        └── flight.go
        └── passenger.go
        └── registration.go
    └── services    
        └── flight_registrator.go
        └── flight_registrator_test.go
    └── storage -- подпакет app, уровень взаимодействия с хранилищами, другими словами адаптер домена к внешним зависимостям, которые отвечают за хранение данных
        ├── postgres -- конкретная реализация 
            ├── registration_repository -- конкретная реализация postgres-репозитория, который отвечает за CRUD для app.Registration 
        └── memory 
    └── utils  
```
