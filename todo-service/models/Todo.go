package models

type Status string

const (
	Pending    Status = "pending"
	InProgress Status = "in-progress"
	Completed  Status = "completed"
)

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Status      Status `json:"status"`
	CreatedBy   string `json:"createdBy"`
	CreatedOn   string `json:"createdOn"`
}
