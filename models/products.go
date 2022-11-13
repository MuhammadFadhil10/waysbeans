package models

type Products struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"productName" form:"productName"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Photo       string `json:"photo" form:"photo"`
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
