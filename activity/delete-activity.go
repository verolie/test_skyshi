package activity

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/verolie/test-skyshi/multi"
)

func DeleteActivity(c echo.Context) error {
	var activity Activities

	id := c.Param("id")
	db := multi.SetDatabase()

	result := db.Delete(&activity, "activity_id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusOK, "failed Get ID")
	}

	return c.JSON(http.StatusOK, multi.ResponseDetail(id))
}
