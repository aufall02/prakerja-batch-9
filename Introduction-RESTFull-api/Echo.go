package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//route handler

	// e.GET("/", HelloController)
	e.POST("/users", CreateUser)
	// e.GET("/user/:id", GetUserByIdController)

	e.Start(":8000")

}

type User struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"nama"`
}

type ResponseCode struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello word")
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	user := User{Id: id, Nama: "Aufal Marom", Email: "aufaluzumaki@gmail.com"}
	return c.JSON(http.StatusOK, map[string]User{
		"user": user,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{Id: id, Nama: "Aufal Marom", Email: "aufaluzumaki@gmail.com"}
	return c.JSON(http.StatusOK, map[string]User{
		"user": user,
	})
}

func UserSearchController(c echo.Context) error {
	nama := c.QueryParam("nama")
	user := User{Id: 1, Nama: nama, Email: "aufaluzumaki@gmail.com"}
	return c.JSON(http.StatusOK, map[string]User{
		"user": user,
	})
}

// func CreateUser(c echo.Context) error {
// 	nama := c.FormValue("nama")
// 	email := c.FormValue("email")
// 	responSucces := ResponseCode{Message: "sucses", Data: User{Nama: nama, Email: email}}

// 	return c.JSON(http.StatusOK, map[string]ResponseCode{
// 		"1" : responSucces,
// 	})
// }

func CreateUser(c echo.Context) error {
	user := User{}
 	c.Bind(&user)

	return c.JSON(http.StatusOK, user)
}