package models

type Task struct {
	ID string `json:"id"`
	Title string `json:"title,required"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
	Status string `json:"status"`
}