package models

type TSc struct {
	StudentID int     `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Score     float64 `json:"score"`
}
