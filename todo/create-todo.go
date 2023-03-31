package todo

import (
	"net/http"
	"time"

	"github.com/verolie/test-skyshi/multi"

	"github.com/verolie/test-skyshi/model"

	"github.com/labstack/echo/v4"
)

type Todos struct {
	Todo_id         int `gorm:"primary_key;auto_increment;not_null"`
	ActicityGroupId int `gorm:"column:activity_group_id"`
	Title           string
	Priority        string
	IsActive        int       `gorm:"column:is_active"`
	CreateAt        time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func CreateUserTodo(c echo.Context) error {
	var active int
	db := multi.SetDatabase()

	ctr := model.TodoCreateRequest{}
	err := c.Bind(&ctr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	IsActive := ctr.IsActive
	if IsActive {
		active = 1
	} else if !IsActive {
		active = 0
	}

	todo := Todos{
		ActicityGroupId: ctr.ActivityGroupId,
		Title:           ctr.Title,
		Priority:        "very-high",
		IsActive:        active,
		CreateAt:        time.Now(),
		UpdatedAt:       time.Now(),
	}

	result := db.Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	resp := lastUpdateTodo(db, 0)
	return c.JSON(http.StatusOK, multi.ResponseDataDetail(resp))
}
