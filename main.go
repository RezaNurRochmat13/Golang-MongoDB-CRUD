package main

import (
	"fmt"
	"log"
	"net/http"
	"svc-users-go/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	UserHandlerPackage "svc-users-go/module/v1/presenter"
	UserRepoPackage "svc-users-go/module/v1/repository"
	UserUseCasePackage "svc-users-go/module/v1/usecase"
)

func main() {
	fmt.Println("Hello world")

	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	echoRouter := echo.New()
	echoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	userRepo := UserRepoPackage.NewUserRepository(mongoConnection)
	userUseCase := UserUseCasePackage.NewUserUseCase(userRepo)
	UserHandlerPackage.NewUserHandler(echoRouter, userUseCase)

	echoRouter.Logger.Fatal(echoRouter.Start(":8081"))
}
