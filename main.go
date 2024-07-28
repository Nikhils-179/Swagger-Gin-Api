package main

import (
	docs "Nikhils-179/go-swagger/docs"
	"Nikhils-179/go-swagger/handler"
	db "Nikhils-179/go-swagger/mongo-db-atlas"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API (User-Details)
// @Version 1
// @Description Sample Description

// @contact.name  Nikhil Shrivastava
// @contact.url https://github.com/Nikhils-179
// @host localhost:8080/ 
// @BasePath /api/v1/user

func main() {

	err := db.ConnectMongoDB()

	if err != nil {
		log.Fatalf("Connection to mongodb failed : %v", err)
	}

	defer func() {
		if err = db.DisconnectMongoDB(); err != nil {
			log.Fatal(err)
		}
	}()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "api/v1/user"
	userAPI := r.Group("/api/v1/user")
	{
		userAPI.GET("/detail/:id", handler.GetUser)
		userAPI.POST("/create", handler.CreateUser)
		userAPI.PUT("/update/:id" , handler.UpdateUser)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
