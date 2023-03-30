package handler

import (
	"fmt"

	"github.com/verolie/test-skyshi/model"

	"gorm.io/gorm"
)

func lastUpdateActivity(db *gorm.DB, row int) interface{} {
	var resp Activities
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

	setResp := model.ActivitiesResponse{
		ID:        resp.Activity_id,
		Title:     resp.Title,
		Email:     resp.Email,
		CreatedAt: resp.CreateAt,
		UpdatedAt: resp.UpdatedAt,
	}

	return setResp
}
