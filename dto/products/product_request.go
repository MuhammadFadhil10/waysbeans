package productdto

type CreateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Photo       string `json:"photo" form:"photo"`
}
