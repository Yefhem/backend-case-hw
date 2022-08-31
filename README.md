# Description

This project that create ticket and purchase ticket is a basic CRUD application. Includes *CREATE* and *GET* operations.

## Features

- Golang
- Postgres
- Gorm 
- Gorilla/mux
- Docker

## Run

Firstly edit `dbConfig.go` 27. line and docker-compose.yml

after that;

`docker-compose build`

`docker-compose up`


# Endpoints and Examples

### **Create Ticket Option**  

*-* http://127.0.0.1:8000/ticket_options

*-* Add new Ticket option with JSON type. Need To Use **POST** method

### Request Body:

```
{
    "name": "Berlin",
    "desc": "lorem ipsum",
    "allocation": 100
}
```

### Response Body:

```
{
    "id": 1,
    "name": "Berlin",
    "desc": "lorem ipsum",
    "allocation": 100
}
```

---------------

### **Get Ticket**  

*-* http://127.0.0.1:8000/ticket/{id}

*-* Get one ticket with the given ID. If there isn't any record returns 404. Need To Use **GET** method

### Request Body:

x

### Response Body:

```
{
    "id": 1,
    "name": "Berlin",
    "desc": "lorem ipsum",
    "allocation": 100
}
```

---------------

### **Purchase from Ticket Option** 

*-* http://127.0.0.1:8000/ticket_options/{id}/purchases

*-* Purchases a quantity of tickets with the given ID. If the transaction is successful returns 200 but if any error or if the purchased ticket is more than the current ticket returns 400. Need To Use **POST** method.

create an Account with email and password

```
http://127.0.0.1:8000/user_create

{
    "user_name": "steve",
    "pass": "123456"
}
```

return response body with includes unique ID

```
{
    "id": "3aa76025-3589-48fc-9997-d384a7c165d5",
    "user_name": "steve",
    "pass": "$2a$10$cTmBIFVQr0rpE0tVLaGWTOlENhUHpwWpjAluOjdTpMsKhniNWjHha"
}
```

request in the following URL and we are going to use *basic auth* so 
you must enter the username and password otherwise return 4xx status code

```
http://127.0.0.1:8000/ticket_options/1/purchases

{
    "quantity": 10,
    "user_id": "3aa76025-3589-48fc-9997-d384a7c165d5"
}
```

---------------

- First Situation:
```
{
    "id": 1,
    "name": "Berlin",
    "desc": "lorem ipsum",
    "allocation": 100
}
```

- After the Transaction:

```
{
    "id": 1,
    "name": "Berlin",
    "desc": "lorem ipsum",
    "allocation": 90
}
```