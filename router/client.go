package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thanapatjitmung/handler"
)

func ClientRouter(e *echo.Echo, h handler.ClientHandler) {
	var jwtKey = []byte("secret-jwt-client")
	clientGroup := e.Group("/client")
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtKey,
		// TokenLookup: "header:x-auth-token",
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	})
	clientGroup.GET("/getdata", h.GetAllData, jwtMiddleware)
	clientGroup.GET("/profile/:id", h.GetProfile, jwtMiddleware)
	clientGroup.PUT("/profile/:id", h.UpdateProfile, jwtMiddleware)
}
