package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_entities "github.com/thanapatjitmung/entities"
	_usecase "github.com/thanapatjitmung/usecase"
)

type (
	AdminHandler interface {
		GetAllData(c echo.Context) error
		GetByIdForAdmin(c echo.Context) error
		UpdateUserForAdmin(c echo.Context) error
		DeleteUserForAdmin(c echo.Context) error
	}

	adminHandlerImpl struct {
		adminUsecase _usecase.AdminUsecase
	}
)

func NewAdminHandlerImpl(adminUsecase _usecase.AdminUsecase) AdminHandler {
	return &adminHandlerImpl{adminUsecase: adminUsecase}
}

func (h adminHandlerImpl) GetAllData(c echo.Context) error {
	data, err := h.adminUsecase.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

func (h *adminHandlerImpl) GetByIdForAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}
	user, err := h.adminUsecase.GetByIdForAdmin(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *adminHandlerImpl) UpdateUserForAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	user := new(_entities.UserRes)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err = h.adminUsecase.UpdateUserForAdmin(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user updated successfully"})
}

func (h *adminHandlerImpl) DeleteUserForAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	err = h.adminUsecase.DeleteUserForAdmin(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user deleted successfully"})
}
