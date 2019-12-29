package presenter

import (
	"log"
	"net/http"
	"strconv"
	"svc-users-go/module/v1/role/model"
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
	groupingPath.POST("/role", injectionHandler.CreateNewRoles)
	groupingPath.GET("/role/:id", injectionHandler.GetDetailRoles)

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

func (rp *RoleHandler) CreateNewRoles(ctx echo.Context) error {
	rolePayload := new(model.CreateRole)

	errorHandlerBindJSON := ctx.Bind(rolePayload)
	if !utils.GlobalErrorException(errorHandlerBindJSON) {
		return errorHandlerBindJSON
	}

	// Save role
	errorHandlerSaveRole := rp.RoleUseCase.CreateNewRole(rolePayload)
	if !utils.GlobalErrorException(errorHandlerSaveRole) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerSaveRole,
			"message": "Error when get usecase",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":      "Role created successfully",
		"created_role": rolePayload,
	})
}

func (rp *RoleHandler) GetDetailRoles(ctx echo.Context) error {
	id := ctx.Param("id")

	// Check param is exist
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Parameter id is required"})
	}

	// Find role by id usecase
	findRoleByIdUseCase, errorHandlerUseCase := rp.RoleUseCase.FindRoleById(id)
	if !utils.GlobalErrorException(errorHandlerUseCase) {
		log.Printf("Error an occured %s", errorHandlerUseCase.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error when get usecase. Check logs for more info"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": findRoleByIdUseCase,
	})
}
