package presenter

import (
	"log"
	"net/http"
	"svc-users-go/module/v1/model"
	"svc-users-go/module/v1/usecase"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUseCase usecase.UseCase
}

func NewUserHandler(e *echo.Echo, userUseCase usecase.UseCase) {
	injectionHandler := &UserHandler{
		UserUseCase: userUseCase,
	}

	groupingPath := e.Group("/api/v1")
	groupingPath.GET("/users", injectionHandler.GetAllUsers)
	groupingPath.GET("/user/:id", injectionHandler.GetDetailUsers)
	groupingPath.POST("/user", injectionHandler.CreateNewUser)
}

func (uh *UserHandler) GetAllUsers(ctx echo.Context) error {
	findAllUserUseCase, errorHandlerUseCase := uh.UserUseCase.FindAllUsers()

	if errorHandlerUseCase != nil {
		log.Println("Error when handle usecase : ", errorHandlerUseCase.Error())

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCase.Error(),
			"message": "Error when get usecase",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"total": len(findAllUserUseCase),
		"count": len(findAllUserUseCase),
		"data":  findAllUserUseCase,
	})
}

func (uh *UserHandler) GetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	findUserById, errorHandlerUseCase := uh.UserUseCase.FindUserById(id)

	if errorHandlerUseCase != nil {
		log.Println("Error when get usecase : ", errorHandlerUseCase.Error())

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": findUserById,
	})
}

func (uh *UserHandler) CreateNewUser(ctx echo.Context) error {
	userPayload := new(model.CreateUser)

	errorHandlerBindJSON := ctx.Bind(userPayload)

	if errorHandlerBindJSON != nil {
		log.Println("Error when bind json : ", errorHandlerBindJSON)
	}

	saveUser, errorHandlerUseCase := uh.UserUseCase.CreateNewUser(userPayload)

	if errorHandlerUseCase != nil {
		log.Println("Error when get usecase : ", errorHandlerUseCase.Error())

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":      "User created successfully",
		"created_user": saveUser,
	})
}
