
# Gin - MongoDBAtlas - Swagger

A REST API exposed to manage data in mongoDB using golang.
Focused on using Swagger.io.

Swagger.io - set of opensource-tools build around OpenAI specification to design , build , document and consume REST API.
- Gin -  WebFramework in Golang
- MongoDB-Atlas - Database as a service 

Packages for Go-lang : 

```bash
swaggerfiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
```







## Table of Content
- Prerequisites
- Folder Structure
- Build and Run the Application
- Accessing the Application
- Configuration Details
- Screeshots


## Prerequisites

Before running the application , ensure you have the following installed:

- Golang 
- Installing MongoDB 7.0 Community Edition(includes mongosh)
- Create Free-Service Account for MongoDB atlas
- go install github.com/swaggo/swag/cmd/swag@latest(install swag init to use)



## Folder Structure
```bash
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── handler
│   └── handler.go
├── main.go
├── models
│   └── models.go
├── mongo-db-atlas
│   └── mongo.go
└── response
    └── response.go

6 directories, 10 files
```
## Build and Run the Application

1. start the mongoDB server . 
run
```bash
mongosh "mongodb+srv://cluster29628.62ifsp3.mongodb.net/?appName=mongosh+2.2.12" --username <username> --password <password>
```

2. generate the swagger documentation in golang
```bash
swag init
```

3. run the server 
```
go run .

```

## Configuration Details

Backend Configuration: The Go server implements the Rest api services  using Gin Framework. It handles the incoming request with designing and tesitng with swagger and writes data in database 

## Screenshot


<img width="800" alt="Screenshot 2024-07-29 at 3 24 47 AM" src="https://github.com/user-attachments/assets/bd06d93c-dc20-4ea0-baba-7943b3de1988">
<img width="800" alt="Screenshot 2024-07-29 at 3 18 56 AM" src="https://github.com/user-attachments/assets/a5ac049d-ca70-47ae-97d2-9635d4a7a829">


