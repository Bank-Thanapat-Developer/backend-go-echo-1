package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_handler "github.com/thanapatjitmung/handler"
	_repository "github.com/thanapatjitmung/repository"
	"github.com/thanapatjitmung/router"
	_usecase "github.com/thanapatjitmung/usecase"
)

func main() {

	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	jwtKey := []byte("secret-jwt-client")
	authRepo := _repository.NewUserRepoImpl("client.csv", "admin.csv")

	authUseCase := _usecase.NewAuthUseCaseImpl(authRepo, jwtKey)

	authHandler := _handler.NewAuthHandlerImpl(authUseCase)

	clientUseCase := _usecase.NewClientUsecaseImpl(authRepo)

	clientHandler := _handler.NewClientHandlerImpl(clientUseCase)

	adminUseCase := _usecase.NewAdminUsecaseImpl(authRepo)

	adminHandler := _handler.NewAdminHandlerImpl(adminUseCase)

	// Routes
	router.AuthRouter(e, authHandler)
	router.ClientRouter(e, clientHandler)
	router.AdminRouter(e, adminHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
