package todo

import (
	"net/http"

	"github.com/verolie/test-skyshi/multi"

	"github.com/labstack/echo/v4"
)

func GetUsersTodo(c echo.Context) error {
	activityId := c.QueryParam("activity_group_id")

	db := multi.SetDatabase()

	var data []Todos

	if activityId != "" {
		result := db.Where("activity_group_id = ?", activityId).Find(&data)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, multi.ResponseDetail(activityId))
		}
	} else {
		result := db.Find(&data)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "empty row")
		}
	}

	return c.JSON(http.StatusOK, multi.ResponseDataDetail(data))
}

func GetUserTodo(c echo.Context) error {
	id := c.Param("id")
	db := multi.SetDatabase()

	var data []Todos

	result := db.Where("todo_id = ?", id).First(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusOK, multi.ResponseDetailTodo(id))
	}

	return c.JSON(http.StatusOK, multi.ResponseDataDetail(data))
}
