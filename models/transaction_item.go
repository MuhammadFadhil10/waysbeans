package models

import "time"

type TransactionItem struct {
	ID        int          `json:"id" gorm:"primaryKey;autoIncrenment"`
	ProductID int          `json:"product_id"`
	Products  Products     `json:"products" gorm:"foreignKey:ProductID;references:ID"`
	UserID    int          `json:"userId" `
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"createdAT" gorm:"default:Now()"`
	UpdateAt  time.Time    `json:"updatedAt" gorm:"default:null"`
}
