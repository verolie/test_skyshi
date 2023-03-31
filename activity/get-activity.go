package activity

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/verolie/test-skyshi/multi"
)

func GetActivities(c echo.Context) error {
	db := multi.SetDatabase()

	var data []Activities

	result := db.Find(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusInternalServerError, "Error retrieving records")
	}

	return c.JSON(http.StatusOK, multi.ResponseDataDetail(data))
}

func GetActivity(c echo.Context) error {
	id := c.Param("id")
	db := multi.SetDatabase()

	var data []Activities

	result := db.Where("activity_id = ?", id).First(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusOK, multi.ResponseDetail(id))
	}

	return c.JSON(http.StatusOK, multi.ResponseDataDetail(data))
}
