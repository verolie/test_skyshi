package todo

import (
	"fmt"

	"github.com/verolie/test-skyshi/model"

	"gorm.io/gorm"
)

func lastUpdateTodo(db *gorm.DB, row int) interface{} {
	var active bool
	var resp Todos
	if row == 0 {
		result := db.Last(&resp)

		if result.Error != nil {
			fmt.Println("gagal")
			return nil
		}
	} else {
		result := db.First(&resp, row)
		if result.Error != nil {
			fmt.Println("gagal")
			return nil
		}
	}

	if resp.IsActive == 1 {
		active = true
	} else if resp.IsActive == 0 {
		active = false
	}

	setResp := model.DotoResponse{
		ID:              resp.Todo_id,
		ActivityGroupID: resp.ActicityGroupId,
		Title:           resp.Title,
		IsActive:        active,
		Priority:        resp.Priority,
		CreatedAt:       resp.CreateAt,
		UpdatedAt:       resp.UpdatedAt,
	}

	return setResp
}
