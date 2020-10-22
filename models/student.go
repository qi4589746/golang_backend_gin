package models

import (
	"time"
)

type Student struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Birthday time.Time `json:"birthday"`
}
