package todo

import (
	"net/http"
	"time"

	"github.com/verolie/test-skyshi/model"

	"github.com/verolie/test-skyshi/handler"

	"github.com/labstack/echo/v4"
)

func UpdateUserTodo(c echo.Context) error {
	var title, priority string
	var active int
	var ctr model.TodoUpdateRequest
	id := c.Param("id")
	db := handler.SetDatabase()

	var data Todos
	result := db.Where("todo_id = ?", id).First(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusOK, handler.ResponseDetailTodo(id))
	}

	err := c.Bind(&ctr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if ctr.Title == "" {
		title = data.Title
	} else {
		title = ctr.Title
	}

	if ctr.Priority == "" {
		priority = data.Priority
	} else {
		priority = ctr.Priority
	}

	IsActive := ctr.IsActive
	if IsActive {
		active = 1
	} else if !IsActive {
		active = 0
	}

	data = Todos{
		Todo_id:         data.Todo_id,
		ActicityGroupId: data.ActicityGroupId,
		Title:           title,
		Priority:        priority,
		IsActive:        active,
		CreateAt:        data.CreateAt,
		UpdatedAt:       time.Now(),
	}

	result = db.Save(&data)
	if result.Error != nil {
		// handle the error
		return c.JSON(http.StatusInternalServerError, "Error updating record")
	}

	resp := lastUpdateTodo(db, data.Todo_id)
	return c.JSON(http.StatusOK, handler.ResponseDataDetail(resp))
}
