# simple-api-for-redis
This project is to practicing set up api that doing CRUD to Redis and swagger is set up as well that client as easliy test the api.

## Basic information:
1. Redis server at localhost and running at port 6379
2. Go version is 1.18.1
3. Gin is running at port 8080, client can visit localhost:8080/swagger/index.html to see all api after running following code at terminal:
```
go run main.go
```

## API description:
1. GET : client can get all existed user data in Redis (no parameter is required)
2. POST : client can add new user information (id, username, password, email)
3. PUT : client can specific user information by id
4. DELETE : client can deleted user information from Redis by id
