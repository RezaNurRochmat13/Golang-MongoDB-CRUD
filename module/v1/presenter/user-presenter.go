package presenter

import (
	"net/http"
	"strconv"
	"svc-users-go/module/v1/model"
	"svc-users-go/module/v1/usecase"
	"svc-users-go/utils"

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
	groupingPath.PUT("/user/:id", injectionHandler.UpdateUser)
	groupingPath.DELETE("/user/:id", injectionHandler.DeleteUser)
}

func (uh *UserHandler) GetAllUsers(ctx echo.Context) error {
	var (
		limitParam      = ctx.QueryParam("limit")
		pagesParam      = ctx.QueryParam("page")
		name            = ctx.QueryParam("name")
		convertLimit, _ = strconv.ParseInt(limitParam, 10, 64)
		convertPage, _  = strconv.ParseInt(pagesParam, 10, 64)
	)

	// Limiting and paging data
	findAllUserUseCase, errorHandlerUseCase := uh.UserUseCase.FindAllUsers(name, convertLimit, convertPage)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCase.Error(),
			"message": "Error when get usecase",
		})
	}

	// Count all data
	countAllDataUser, errorHandlerUseCaseCount := uh.UserUseCase.CountAllUsers()
	if !utils.GlobalErrorException(errorHandlerUseCaseCount) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCaseCount.Error(),
			"message": "Error when get usecase",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllUserUseCase),
		"data":  findAllUserUseCase,
		"total": countAllDataUser,
		"limit": convertLimit,
		"page":  convertPage,
	})
}

func (uh *UserHandler) GetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	findUserById, errorHandlerUseCase := uh.UserUseCase.FindUserById(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {

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

	if !utils.GlobalErrorDatabaseException(errorHandlerBindJSON) {
		return errorHandlerBindJSON
	}

	saveUser, errorHandlerUseCase := uh.UserUseCase.CreateNewUser(userPayload)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":      "User created successfully",
		"created_user": saveUser,
	})
}

func (uh *UserHandler) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userUpdate := new(model.UpdateUser)

	errorHandlerBindJSON := ctx.Bind(userUpdate)

	if !utils.GlobalErrorDatabaseException(errorHandlerBindJSON) {
		return errorHandlerBindJSON
	}

	errorHandlerUpdate := uh.UserUseCase.UpdateUser(id, userUpdate)

	if !utils.GlobalErrorDatabaseException(errorHandlerUpdate) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Update gagal",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "User updated sucessfully"})
}

func (uh *UserHandler) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	errorHandlerDelete := uh.UserUseCase.DeleteUser(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerDelete) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "User deleted sucessfully"})
}
