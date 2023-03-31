package activity

import (
	"net/http"
	"time"

	"github.com/verolie/test-skyshi/model"
	"github.com/verolie/test-skyshi/multi"

	"github.com/labstack/echo/v4"
)

type Activities struct {
	Activity_id int `gorm:"primary_key;auto_increment;not_null"`
	Title       string
	Email       string
	CreateAt    time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func CreateActivity(c echo.Context) error {

	db := multi.SetDatabase()

	ctr := model.CreateUsersRequest{}
	err := c.Bind(&ctr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	activities := Activities{
		Title:     ctr.Title,
		Email:     ctr.Email,
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.Create(&activities)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	resp := lastUpdateActivity(db, 0)

	return c.JSON(http.StatusOK, multi.ResponseDataDetail(resp))
}
