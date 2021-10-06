package controllers

import (
	"echo-rest/helpers"
	"echo-rest/models"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CheckLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.CheckLogin(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token, := jwt.New() 
	// return c.String(http.StatusOK, "Loginn Success")
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
