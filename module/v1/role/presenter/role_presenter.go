package presenter

import (
	"net/http"
	"strconv"
	"svc-users-go/module/v1/role/usecase"
	"svc-users-go/utils"

	"github.com/labstack/echo"
)

type RoleHandler struct {
	RoleUseCase usecase.UseCase
}

func NewRoleHandler(e *echo.Echo, roleUseCase usecase.UseCase) {
	injectionHandler := &RoleHandler{
		RoleUseCase: roleUseCase,
	}

	groupingPath := e.Group("/api/v1")
	groupingPath.GET("/roles", injectionHandler.GetAllRoles)

}

func (rp *RoleHandler) GetAllRoles(ctx echo.Context) error {
	var (
		limitParam      = ctx.QueryParam("limit")
		pagesParam      = ctx.QueryParam("page")
		convertLimit, _ = strconv.ParseInt(limitParam, 10, 64)
		convertPage, _  = strconv.ParseInt(pagesParam, 10, 64)
	)

	// Find all roles
	findAllRolesUseCase, errorHandlerUseCase := rp.RoleUseCase.FindAllRoles(convertLimit, convertPage)
	if !utils.GlobalErrorException(errorHandlerUseCase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCase.Error(),
			"message": "Error when get usecase",
		})
	}

	// Count all roles
	countAllRoles, errorHandlerCount := rp.RoleUseCase.CountAllRoles()
	if !utils.GlobalErrorException(errorHandlerCount) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerCount.Error(),
			"message": "Error when get usecase",
		})
	}

	// Checking data exist or not
	if findAllRolesUseCase == nil {
		return ctx.JSON(http.StatusOK, echo.Map{
			"count": len(findAllRolesUseCase),
			"data":  "Data kosong",
			"total": countAllRoles,
			"limit": convertLimit,
			"page":  convertPage,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllRolesUseCase),
		"data":  findAllRolesUseCase,
		"total": countAllRoles,
		"limit": convertLimit,
		"page":  convertPage,
	})
}
