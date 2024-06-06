package models

type Task struct {
	ID     int    `json:"id",omitempty`
	Name   string `json:"name"`
	IsDone bool   `json:"isDone"`
}

type TaskResponse struct {
	ID     int    `json:"id", omitempty`
	Name   string `json:"name", omitempty`
	IsDone bool   `json:"name", omitempty`
}
