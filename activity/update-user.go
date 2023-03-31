package activity

import (
	"net/http"
	"time"

	"github.com/verolie/test-skyshi/model"
	"github.com/verolie/test-skyshi/multi"

	"github.com/labstack/echo/v4"
)

func UpdateActivity(c echo.Context) error {
	var ctr model.UpdateTitleRequest
	id := c.Param("id")
	db := multi.SetDatabase()

	var data Activities

	result := db.Where("activity_id = ?", id).First(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusOK, multi.ResponseDetail(id))
	}

	err := c.Bind(&ctr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	data = Activities{
		Activity_id: data.Activity_id,
		Title:       ctr.Title,
		Email:       data.Email,
		CreateAt:    data.CreateAt,
		UpdatedAt:   time.Now(),
	}

	result = db.Save(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusInternalServerError, "Error updating record")
	}

	resp := lastUpdateActivity(db, data.Activity_id)
	return c.JSON(http.StatusOK, multi.ResponseDataDetail(resp))
}
