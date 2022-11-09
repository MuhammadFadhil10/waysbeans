package models

type Products struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Photo       string `json:"photo"`
}

type ProductTransactionResponse struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Description   string `json:"description"`
	Stock         int    `json:"stock"`
	Photo         string `json:"photo"`
	OrderQuantity int    `json:"orderQuantity"`
}

func (ProductTransactionResponse) TableName() string {
	return "products"
}
