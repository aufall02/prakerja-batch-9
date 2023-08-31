package routes

import (
	"dborm/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:nama", controllers.GetUserController)
	e.DELETE("/users/:nama", controllers.DeleteUserController)
	e.PUT("/users/:nama", controllers.UpdateUserController)

	return e
}