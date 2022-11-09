package models

import "time"

type Cart struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Qty        int          `json:"qty"`
	TotalPrice int          `json:"totalPrice"`
	ProductID  int          `json:"productId"`
	Products   Products     `json:"products" gorm:"foreignKey:ProductID;references:ID"`
	UserID     int          `json:"userId" `
	User       UserResponse `json:"user" `
	CreatedAt  time.Time    `json:"createdAT" gorm:"default:Now()"`
	UpdateAt   time.Time    `json:"updatedAt" gorm:"default:null"`
}

type CartUpdateRequest struct {
	ID         int `json:"id"`
	Qty        int `json:"qty"`
	TotalPrice int `json:"totalPrice"`
}

func (CartUpdateRequest) TableName() string {
	return "carts"
}
