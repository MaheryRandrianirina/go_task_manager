package utils

type Task struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
