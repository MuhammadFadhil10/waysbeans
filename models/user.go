package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement" `
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
