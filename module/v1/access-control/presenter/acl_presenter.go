package presenter

import (
	"net/http"
	"strconv"
	"svc-users-go/module/v1/access-control/usecase"
	"svc-users-go/utils"

	"github.com/labstack/echo"
)

type AccessControlHandler struct {
	AccessControlUseCase usecase.UseCase
}

func NewAccessControlHandler(e *echo.Echo, accessControlUseCase usecase.UseCase) {
	injectionHandler := &AccessControlHandler{
		AccessControlUseCase: accessControlUseCase,
	}
	groupingPath := e.Group("/api/v1")
	groupingPath.GET("/access-controls", injectionHandler.GetAllAccessControl)
	groupingPath.GET("/access-control/:id", injectionHandler.GetDetailAccessControl)

}

func (ap *AccessControlHandler) GetAllAccessControl(ctx echo.Context) error {
	var (
		limitParam      = ctx.QueryParam("limit")
		pagesParam      = ctx.QueryParam("page")
		convertLimit, _ = strconv.ParseInt(limitParam, 10, 64)
		convertPage, _  = strconv.ParseInt(pagesParam, 10, 64)
	)
	// Count all access control
	countAllUserUseCase, errorHandlerUsecase := ap.AccessControlUseCase.CountAllUser()
	if !utils.GlobalErrorDatabaseException(errorHandlerUsecase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errorHandlerUsecase.Error(),
		})
	}

	// Find all access control
	findAllAccessControlUsecase, errorHandlerUsecase := ap.AccessControlUseCase.FindAllUser(convertLimit, convertPage)
	if !utils.GlobalErrorException(errorHandlerUsecase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errorHandlerUsecase.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllAccessControlUsecase),
		"total": countAllUserUseCase,
		"data":  findAllAccessControlUsecase,
		"page":  convertPage,
		"limit": convertLimit,
	})
}

func (ap *AccessControlHandler) GetDetailAccessControl(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusOK, echo.Map{"message": "Parameter is required"})
	}

	// Find access control by id
	findAccessControlById, errorHandlerUseCase := ap.AccessControlUseCase.FindAccessControlById(id)
	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": errorHandlerUseCase.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": findAccessControlById,
	})
}
