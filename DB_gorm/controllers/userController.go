package controllers

import (
	"dborm/lib/database"
	"dborm/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {

	namaUser := c.Param("nama")
	user, err := database.GetUser(namaUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    user,
	})

}

func CreateUserController(c echo.Context) error {

	var user = models.User{}
	c.Bind(&user)
	

	userCreate, err := database.CreateUser(user)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    userCreate,
	})

}

func DeleteUserController(c echo.Context) error {

	namaUser := c.Param("nama")
	_, err := database.DeleteUser(namaUser)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
	})

}

func UpdateUserController(c echo.Context) error {

	var user = models.User{}
	c.Bind(&user)
	namaUser := c.Param("nama")

	userCreate, err := database.UpdateUser(user, namaUser)
	if err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succses",
		"data":    userCreate,
	})

}