package multi

import (
	"github.com/verolie/test-skyshi/model"
)

func ResponseDataDetail(data interface{}) model.DataStatusResponse {
	resp := model.DataStatusResponse{
		Status:  "Success",
		Message: "Success",
		Data:    data,
	}
	return resp
}

func ResponseDetail(id string) model.StatusResponse {
	resp := model.StatusResponse{
		Status:  "Not Found",
		Message: "Activity with ID " + id + " Not Found",
	}
	return resp
}

func ResponseDetailActivityDelete(id string) model.StatusResponse {
	resp := model.StatusResponse{
		Status:  "Not Found",
		Message: "Activity with ID " + id + " Not Found or it have connection to todo table",
	}
	return resp
}

func ResponseDetailTodo(id string) model.StatusResponse {
	resp := model.StatusResponse{
		Status:  "Not Found",
		Message: "Todo with ID " + id + " Not Found",
	}
	return resp
}
