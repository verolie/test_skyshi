package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	var activity Activities

	id := c.Param("id")
	db := SetDatabase()

	result := db.Delete(&activity, id)
	if result.Error != nil {
		return c.JSON(http.StatusOK, ResponseDetailActivityDelete(id))
	}

	return c.JSON(http.StatusOK, ResponseDetail(id))
}
