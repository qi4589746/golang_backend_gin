package models

type Class struct {
	ID         uint   `json:"id"`
	Name       int    `json:"name"`
	Subject    string `json:"subject"`
	CreateYear int    `json:"create_year"`
	TeacherId  uint   `json:"teacher_id"`
}
