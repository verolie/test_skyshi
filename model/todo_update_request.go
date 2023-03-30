package model

type TodoUpdateRequest struct {
	Title    string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool   `json:"is_active"`
	Status   string `json:"status"`
}
