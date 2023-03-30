package model

type TodoCreateRequest struct {
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
}
