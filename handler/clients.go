package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_entities "github.com/thanapatjitmung/entities"
	_usecase "github.com/thanapatjitmung/usecase"
)

type (
	ClientHandler interface {
		GetAllData(c echo.Context) error
		GetProfile(c echo.Context) error
		UpdateProfile(c echo.Context) error
	}

	clientHandlerImpl struct {
		clientUsecase _usecase.ClientUsecase
	}
)

func NewClientHandlerImpl(clientUsecase _usecase.ClientUsecase) ClientHandler {
	return &clientHandlerImpl{
		clientUsecase: clientUsecase,
	}
}

func (h *clientHandlerImpl) GetAllData(c echo.Context) error {
	data, err := h.clientUsecase.GetAllDataClient()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	// fmt.Println(data)
	return c.JSON(http.StatusOK, data)
}
func (h *clientHandlerImpl) GetProfile(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}
	user, err := h.clientUsecase.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *clientHandlerImpl) UpdateProfile(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	user := new(_entities.UserRes)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err = h.clientUsecase.UpdateProfile(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "user updated successfully"})
}
