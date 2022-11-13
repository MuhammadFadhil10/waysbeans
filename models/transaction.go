package models

import "time"

type Transaction struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int          `json:"userId"`
	User       UserResponse `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Status     string       `json:"status"`
	Products   []Products   `json:"products" gorm:"many2many:transaction_products;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TotalPrice int          `json:"totalPrice"`
	CreatedAt  time.Time    `json:"createdAT" gorm:"default:Now()"`
	UpdateAt   time.Time    `json:"updatedAt" gorm:"default:null"`
}

// gorm:"foreignKey:UserID;references:ID"
