package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func TakeAdmin(c echo.Context) error {
	result, err := models.TakeAdmin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AllAdmin(c echo.Context) error {
	nameadm := c.FormValue("nameadm")
	jk := c.FormValue("jk")
	telp := c.FormValue("telp")

	result, err := models.AllAdmin(nameadm, jk, telp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateAdmin(c echo.Context) error {
	idadm := c.FormValue("idadm")
	nameadm := c.FormValue("nameadm")
	jk := c.FormValue("jk")
	telp := c.FormValue("telp")

	conv_id, err := strconv.Atoi(idadm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateAdmin(conv_id, nameadm, jk, telp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteAdmin(c echo.Context) error {
	idadm := c.FormValue("idadm")

	conv_id, err := strconv.Atoi(idadm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteAdmin(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
