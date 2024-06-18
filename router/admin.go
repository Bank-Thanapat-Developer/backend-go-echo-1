package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thanapatjitmung/handler"
)

func AdminRouter(e *echo.Echo, h handler.AdminHandler) {
	var jwtKey = []byte("secret-jwt-admin")

	adminGroup := e.Group("/admin")
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  jwtKey,
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	})

	adminGroup.GET("/getdata", h.GetAllData, jwtMiddleware)
	adminGroup.GET("/getdata/:id", h.GetByIdForAdmin, jwtMiddleware)
	adminGroup.PUT("/users/:id", h.UpdateUserForAdmin, jwtMiddleware)
	adminGroup.DELETE("/users/:id", h.DeleteUserForAdmin, jwtMiddleware)
}
