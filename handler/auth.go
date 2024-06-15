package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_entities "github.com/thanapatjitmung/entities"
	_usecase "github.com/thanapatjitmung/usecase"
)

type (
	AuthHandler interface {
		Register(c echo.Context) error
		Login(c echo.Context) error
	}

	authHandlerImpl struct {
		authUsecase _usecase.AuthUsecase
	}
)

func NewAuthHandlerImpl(authUsecase _usecase.AuthUsecase) AuthHandler {
	return &authHandlerImpl{authUsecase: authUsecase}
}

func (a *authHandlerImpl) Register(c echo.Context) error {
	user := new(_entities.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	err = a.authUsecase.Register(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Register Successfully"})
}

func (a *authHandlerImpl) Login(c echo.Context) error {
	user := new(_entities.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	token, err := a.authUsecase.Login(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
