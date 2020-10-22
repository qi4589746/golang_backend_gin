package models

type Teacher struct {
	ID      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Gender  string `json:"gender" form:"name"`
	Subject string `json:"subject" form:"subject"`
}
