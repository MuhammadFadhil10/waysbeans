package models

type Transaction struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int          `json:"userId"`
	User       UserResponse `json:"user"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Address    string       `json:"address"`
	Attachment string       `json:"attachment"`
	Status     string       `json:"status"`
	ProductID  int          `json:"productId"`
	Products   Products     `json:"products" gorm:"foreignKey:ProductID;references:ID"`
}

// gorm:"foreignKey:UserID;references:ID"
