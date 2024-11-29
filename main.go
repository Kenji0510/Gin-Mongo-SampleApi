package main

import (
	"context"
	"fmt"
	"gin-api-sample/config"
	"gin-api-sample/controller"
	domain "gin-api-sample/domain/users/usecase"
	"gin-api-sample/routes"
	"gin-api-sample/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	us          domain.UserUsecaseI
	uc          routes.UserController
	ctx         context.Context
	mongoClient *mongo.Client
)

func init() {
	ctx = context.TODO()

	// Mongo
	mongoCon, err := config.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("mongo connection established")
	ur := service.NewUserRepository(mongoCon)
	us = controller.NewUserUsecase(ur, ctx)
	uc = routes.New(us)
	server = gin.Default()
}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, ctx)

	basePath := server.Group("/v1")
	uc.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))

}
