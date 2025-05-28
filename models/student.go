package models

// type Student struct {
// 	ID    uint   `gorm:"primaryKey" json:"id"`
// 	Name  string `json:"name"`
// 	Email string `gorm:"unique" json:"email"`
// 	Phone string `json:"phone"`
// }

type Student struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone"`
	Password string `json:"password"` // Hide password in JSON responses
}
