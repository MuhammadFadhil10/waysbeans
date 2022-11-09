package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement" `
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
