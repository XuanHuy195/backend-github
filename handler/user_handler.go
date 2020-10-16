package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user": "Huy",
		"email": "huy@gmail.com",
	})
}
func HandleSignUp(c echo.Context) error {
	type User struct {
		Email string `json:"email"`
		Name string `json:"name"`
		Age int `json:"age"`
	}
	user := User {
		Email: "Trần Xuân",
		Name: "Huy",
		Age: 20,
	}
	return c.JSON(http.StatusOK, user)
}
