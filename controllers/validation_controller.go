package controllers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

//testValidation
type Customer struct {
	Name    string `validate: "required"`
	Email   string `validate: "required,email"`
	Address string `validate: "required"`
	Age     int    `validate: "gte=18,lte=40"`
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "saroh@gmail.com"
	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Not valid",
		})

	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Name:    "Sarohh",
		Email:   "saroh@gmail.com",
		Address: "",
		Age:     19,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
