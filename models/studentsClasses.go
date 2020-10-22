package models

type StudentsClasses struct {
	ID        uint `json:"id"`
	StudentId uint `json:"student_id"`
	ClassId   uint `json:"class_id"`
}
