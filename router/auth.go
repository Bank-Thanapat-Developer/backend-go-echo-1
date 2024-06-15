package router

import (
	"github.com/labstack/echo/v4"
	"github.com/thanapatjitmung/handler"
)


func AuthRouter(e *echo.Echo, h handler.AuthHandler) {
	authGroup := e.Group("/auth")

	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)

}
