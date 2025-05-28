package models

type Course struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Code  string `gorm:"unique" json:"code"`
}
