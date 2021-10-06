package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func TakeProduct(c echo.Context) error {
	result, err := models.TakeProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AllProduct(c echo.Context) error {
	productname := c.FormValue("productname")
	category := c.FormValue("category")
	image := c.FormValue("image")
	tutorial := c.FormValue("tutorial")

	result, err := models.AllProduct(productname, category, image, tutorial)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	idproduct := c.FormValue("idproduct")
	productname := c.FormValue("productname")
	category := c.FormValue("category")
	image := c.FormValue("image")
	tutorial := c.FormValue("tutorial")

	conv_id, err := strconv.Atoi(idproduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateProduct(conv_id, productname, category, image, tutorial)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteProduct(c echo.Context) error {
	idproduct := c.FormValue("idproduct")

	conv_id, err := strconv.Atoi(idproduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteProduct(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
