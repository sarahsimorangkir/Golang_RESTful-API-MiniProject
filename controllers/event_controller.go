package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func TakeEvent(c echo.Context) error {
	result, err := models.TakeEvent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AllEvent(c echo.Context) error {
	eventname := c.FormValue("eventname")
	eventschedule := c.FormValue("eventschedule")
	eventdescription := c.FormValue("eventdescription")
	image := c.FormValue("image")

	result, err := models.AllEvent(eventname, eventschedule, eventdescription, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateEvent(c echo.Context) error {
	idevent := c.FormValue("idevent")
	eventname := c.FormValue("eventname")
	eventschedule := c.FormValue("eventschedule")
	eventdescription := c.FormValue("eventdescription")
	image := c.FormValue("image")

	conv_id, err := strconv.Atoi(idevent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateEvent(conv_id, eventname, eventschedule, eventdescription, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteEvent(c echo.Context) error {
	idevent := c.FormValue("idevent")

	conv_id, err := strconv.Atoi(idevent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEvent(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
