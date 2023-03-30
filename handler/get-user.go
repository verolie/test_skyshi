package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := SetDatabase()

	var data []Activities

	result := db.Find(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusInternalServerError, "Error retrieving records")
	}

	return c.JSON(http.StatusOK, ResponseDataDetail(data))
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	db := SetDatabase()

	var data []Activities

	result := db.Where("activity_id = ?", id).First(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusOK, ResponseDetail(id))
	}

	return c.JSON(http.StatusOK, ResponseDataDetail(data))
}
