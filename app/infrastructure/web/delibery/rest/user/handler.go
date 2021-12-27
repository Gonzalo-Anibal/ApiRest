package user

import (
	"github.com/Timming/app/application"
	"github.com/labstack/echo"
)

type userHandler struct {
	useCase application.User
}

func NewUserHandler(e *echo.Echo, useCase application.User) *userHandler{
	u := &userHandler{
		useCase: useCase,
	}


	return u
}