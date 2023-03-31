package todo

import (
	"errors"
	"net/http"

	"github.com/verolie/test-skyshi/multi"

	"github.com/verolie/test-skyshi/model"

	"github.com/labstack/echo/v4"
)

func DeleteTodo(c echo.Context) error {
	var activity Todos

	ctr := model.UpdateTitleRequest{}
	err := c.Bind(&ctr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	id := c.Param("id")
	db := multi.SetDatabase()

	result := db.Where("title = ?", ctr.Title).Where("todo_id = ?", id).Delete(&activity)
	if result.Error != nil {
		return c.JSON(http.StatusOK, multi.ResponseDetail(id+" or title"))
	}
	if result.RowsAffected == 0 {
		return errors.New("no records were deleted")
	}

	return c.JSON(http.StatusOK, multi.ResponseDetailTodo(id))
}
