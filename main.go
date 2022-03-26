package main

import (
	"log"
	"triadmoko-be-golang/config"
	"triadmoko-be-golang/handler"
	"triadmoko-be-golang/repository"
	"triadmoko-be-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db := config.Database()

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := service.NewServiceUser(repositoryUser)
	handler := handler.NewHandlerUser(serviceUser)

	router := gin.Default()

	user := router.Group("/api/v1/user")
	user.GET("/", handler.GetUser)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
