package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement" `
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
	Role     string `json:"role" gorm:"default:user"`
}

type UserLoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Photo string `json:"photo"`
	Role  string `json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement" `
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Role     string `json:"role"`
}

type CheckAuthResponse struct {
	Token string `json:"token"`
}

func (UserResponse) TableName() string {
	return "users"
}
func (UserLoginResponse) TableName() string {
	return "users"
}

func (CheckAuthResponse) TableName() string {
	return "users"
}
