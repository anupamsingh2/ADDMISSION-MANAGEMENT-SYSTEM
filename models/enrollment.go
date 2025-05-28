package models

type Enrollment struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}
